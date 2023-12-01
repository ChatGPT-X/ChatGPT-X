package e

import "chatgpt_x/pkg/logger"

var MsgFlags = map[int]string{
	SUCCESS:                  "success",
	ERROR:                    "fail",
	InvalidParams:            "请求参数错误",
	ErrorGetBookDetailFail:   "获取书籍详情失败",
	ErrorGetAuthorDetailFail: "获取作者详情失败",
	ErrorGetBookContentFail:  "获取书籍内容失败",
	ErrorGetBookListFail:     "获取书籍列表失败",
	ErrorGetCategoryListFail: "获取分类列表失败",
	ErrorTranslateFail:       "翻译失败",
	ErrorTranslateInProgress: "正在翻译中，请稍后重试",
	ErrorBookNotExist:        "书籍不存在",
}

// GetMsg get error information based on Code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// HasError any errors will be saved to the log.
func HasError(err error) bool {
	if err != nil {
		sugar := logger.Logger.Sugar()
		sugar.Errorf("An unpredictable error was caught: %s", err)
		return true
	}
	return false
}
