package app

import (
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// Response 基础返回的数据格式。
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseDataList 返回 List 专用的数据格式。
type ResponseDataList struct {
	List      interface{} `json:"list"`
	Page      int64       `json:"page"`       // 当前页数
	PageSize  int64       `json:"page_size"`  // 每页的数据条数
	PageCount int64       `json:"page_count"` // 总共有多少页
	Count     int64       `json:"count"`      // 总共有多少条数据
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
