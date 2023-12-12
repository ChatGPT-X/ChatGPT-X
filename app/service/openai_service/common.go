package openai_service

import (
	"bytes"
	"chatgpt_x/app/models/ai_model"
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/app/models/user"
	"chatgpt_x/pkg/logger"
	"context"
	"fmt"
	"github.com/imroc/req/v3"
	"io"
	"time"
)

var SystemSetting map[string]any

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

// SendRequest 发送请求。
func SendRequest(reqType, method, url string, headers map[string]string, body any) (string, error) {
	client := req.C()
	switch reqType {
	case "web":
		client = client.SetBaseURL(SystemSetting["WebBaseUrl"].(string))
		client = client.SetTimeout(time.Duration(SystemSetting["WebTimeout"].(uint)))
	case "api":
		client = client.SetBaseURL(SystemSetting["ApiBaseUrl"].(string))
		client = client.SetTimeout(time.Duration(SystemSetting["ApiTimeout"].(uint)))
	default:
		panic("invalid request type")
	}
	request := client.R().SetContext(context.Background())
	request = request.SetHeaders(headers)
	request = request.SetBody(body)
	resp, err := request.Send(method, url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.String(), nil
}

// SendStreamRequest 发送流式请求。
func SendStreamRequest(reqType, method, url string, headers map[string]string, body any) (<-chan []byte, error) {
	client := req.C()
	switch reqType {
	case "web":
		client = client.SetBaseURL(SystemSetting["WebBaseUrl"].(string))
		client = client.SetTimeout(time.Duration(SystemSetting["WebTimeout"].(uint)) * time.Second)
	case "api":
		client = client.SetBaseURL(SystemSetting["ApiBaseUrl"].(string))
		client = client.SetTimeout(time.Duration(SystemSetting["ApiTimeout"].(uint)) * time.Second)
	default:
		panic("invalid request type")
	}
	request := client.R().SetContext(context.Background())
	request = request.SetHeaders(headers)
	request = request.SetBody(body)
	resp, err := request.Send(method, url)
	if err != nil {
		return nil, err
	}
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		reader := resp.Response.Body
		defer reader.Close()
		var buffer bytes.Buffer
		for {
			buf := make([]byte, 1) // 1 byte per read
			n, err := reader.Read(buf)
			if err == io.EOF {
				if buffer.Len() > 0 {
					//fmt.Println(buffer.String())
					ch <- buffer.Bytes()
				}
				break
			}
			if err != nil {
				logger.Error("read response body error: ", err)
				break
			}
			if buf[0] == '\n' {
				if buffer.Len() > 0 {
					//fmt.Println(buffer.String())
					ch <- buffer.Bytes()
				}
				buffer.Reset()
				continue
			}
			buffer.Write(buf[:n])
		}
	}()
	return ch, nil
}
