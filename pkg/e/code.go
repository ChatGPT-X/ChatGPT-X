package e

const (
	SUCCESS                  = 200   // 成功
	ERROR                    = 500   // 错误
	InvalidParams            = 400   // 参数错误
	ErrorGetBookDetailFail   = 10001 // 获取书籍详情失败
	ErrorGetAuthorDetailFail = 10002 // 获取作者详情失败
	ErrorGetBookContentFail  = 10003 // 获取书籍内容失败
	ErrorGetBookListFail     = 10004 // 获取书籍列表失败
	ErrorGetCategoryListFail = 10005 // 获取分类列表失败
	ErrorTranslateFail       = 10006 // 翻译失败
	ErrorTranslateInProgress = 10007 // 正在进行翻译，请稍后重试
	ErrorBookNotExist        = 10008 // 书籍不存在
)
