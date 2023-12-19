package openai_service

import (
	"bufio"
	"bytes"
	"chatgpt_x/app/models/ai_model"
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/app/models/user"
	"chatgpt_x/app/service"
	"chatgpt_x/pkg/logger"
	rds "chatgpt_x/pkg/redis"
	"context"
	"fmt"
	"github.com/imroc/req/v3"
	"io"
	"net/http"
	"net/url"
	"time"
)

var ctx = context.Background()

// GetBasicHeaders 获取基础请求头。
func GetBasicHeaders(aiToken string, isEventStream bool) map[string]string {
	headers := map[string]string{
		"Authorization":   "Bearer " + aiToken,
		"Content-Type":    "application/json; charset=utf-8",
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Referer":         "https://chat.openai.com",
		"Origin":          "https://chat.openai.com",
		"Cache-Control":   "no-cache",
		"Pragma":          "no-cache",
		"Accept-Language": "en-US",
	}
	if isEventStream {
		headers["Accept"] = "text/event-stream"
	}
	return headers
}

// GetAiTokenFromUser 根据用户 ID 获取 AI 密钥。
func GetAiTokenFromUser(userID uint) (ai_token.AiToken, error) {
	// 获取用户信息
	userModel, err := user.Get(userID)
	if err != nil {
		return ai_token.AiToken{}, err
	}
	// 检查用户是否被禁用
	if userModel.Status == user.StatusDisable {
		return ai_token.AiToken{}, fmt.Errorf("user is disable: %s", userModel.Username)
	}
	// 获取 AI 密钥信息
	aiTokenModel, err := ai_token.Get(*userModel.AiTokenID)
	if err != nil {
		return ai_token.AiToken{}, err
	}
	// 检查 AI 密钥是否被禁用
	if aiTokenModel.Status == ai_model.StatusDisable {
		return ai_token.AiToken{}, fmt.Errorf("ai token is disable: %s", aiTokenModel.Token)
	}
	return aiTokenModel, nil
}

// clientSetting 设置客户端（基础地址、代理、超时时间等）。
func clientSetting(reqType string, client *req.Client) (*req.Client, error) {
	rdb := rds.RDB
	var baseurl, proxy, timeout string
	switch reqType {
	case "web":
		baseurl = service.RedisSettingOpenaiWebBaseUrl
		proxy = service.RedisSettingOpenaiWebProxy
		timeout = service.RedisSettingOpenaiWebTimeout
	case "api":
		baseurl = service.RedisSettingOpenaiApiBaseUrl
		proxy = service.RedisSettingOpenaiApiProxy
		timeout = service.RedisSettingOpenaiApiTimeout
	default:
		return nil, fmt.Errorf("invalid request type: %s", reqType)
	}
	// 设置基础 URL
	urlVal, err := rdb.Get(ctx, baseurl).Result()
	if err != nil {
		return nil, err
	}
	client = client.SetBaseURL(urlVal)
	// 设置代理
	proxyVal, err := rdb.Get(ctx, proxy).Result()
	if err != nil {
		return nil, err
	}
	client = client.SetProxy(func(request *http.Request) (*url.URL, error) {
		// 注意！这里为空的时候不要去设置代理
		// 否则报 tcp: dial tcp :0: connect: can't assign requested address 错误
		if proxyVal == "" {
			return nil, nil
		}
		return url.Parse(proxyVal)
	})

	// 设置超时时间
	val, err := rdb.Get(ctx, timeout).Int()
	if err != nil {
		return nil, err
	}
	client = client.SetTimeout(time.Duration(val) * time.Second)
	return client, nil
}

// SendRequest 发送常规请求。
func SendRequest(reqType, method, url string, headers map[string]string, body any) (*ResponseResult, error) {
	client := req.C()
	client, err := clientSetting(reqType, client)
	if err != nil {
		return nil, err
	}
	resp, err := client.R().
		SetContext(context.Background()).
		SetHeaders(headers).
		SetBody(body).
		Send(method, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseResult := &ResponseResult{
		StatusCode: resp.StatusCode,
		BodyType:   resp.Header.Get("Content-Type"),
		BodyStream: nil,
		Body:       resp.Bytes(),
	}
	return responseResult, nil
}

// ResponseResult 包含响应的状态码、响应体类型和响应体。
type ResponseResult struct {
	StatusCode int
	BodyType   string
	BodyStream chan []byte // 用于流式响应
	Body       []byte      // 用于普通响应
}

// SendStreamRequest 发送流式请求。
// 修改后的函数现在返回一个包含状态码、响应体类型和响应体的结构体。
func SendStreamRequest(reqType, method, url string, headers map[string]string, body any) (*ResponseResult, error) {
	client := req.C()
	client, err := clientSetting(reqType, client)
	if err != nil {
		return nil, err
	}
	resp, err := client.R().
		SetContext(context.Background()).
		SetHeaders(headers).
		SetBody(body).
		Send(method, url)
	if err != nil {
		return nil, err
	}

	responseResult := &ResponseResult{
		StatusCode: resp.StatusCode,
		BodyType:   resp.Header.Get("Content-Type"),
		BodyStream: make(chan []byte),
		Body:       nil,
	}
	// 根据响应类型处理响应体
	switch resp.Header.Get("Content-Type") {
	case "application/json", "application/json; charset=utf-8":
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		responseResult.Body = bodyBytes
		responseResult.BodyStream = nil
	case "text/event-stream", "text/event-stream; charset=utf-8":
		go handleStreamResponse(resp, responseResult.BodyStream)
	}
	return responseResult, nil
}

// handleStreamResponse 处理流式响应。
func handleStreamResponse(resp *req.Response, bodyStream chan []byte) {
	defer close(bodyStream)
	reader := bufio.NewReaderSize(resp.Response.Body, 2048)
	defer resp.Body.Close()
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("read response body error: ", err)
			return
		}
		// 修正数据后发送
		data = bytes.TrimLeft(data, "data: ")
		data = bytes.TrimRight(data, "\n")
		if len(data) > 0 {
			bodyStream <- data
		}
	}
}
