package app

import (
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// Response 返回的数据格式.
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PageDetail 分页详情.
type PageDetail struct {
	Page         int
	PageSize     int
	TotalPages   int
	TotalResults int
}

// ResponseDataList 返回 List 专用的数据格式.
type ResponseDataList struct {
	List      interface{} `json:"list"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	PageCount int         `json:"page_count"`
	Count     int         `json:"count"`
}

// Response 设置 gin.JSON.
func (g *Gin) Response(httpCode, errCode int, err error, data interface{}) {
	if err != nil {
		logger.Error("Response: httpCode: %d, errCode: %d, err: %v, data: %v\n", httpCode, errCode, err, data)
	}
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
