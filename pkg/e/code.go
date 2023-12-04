package e

const (
	SUCCESS       = 200 // 成功
	ERROR         = 500 // 系统内部未知错误
	InvalidParams = 400 // 请求参数错误
	ErrorAuthFail = 401 //	认证失败或权限不足
)
const (
	ErrorUserIsExist                 = 10001 + iota // 用户已经存在
	ErrorUserCreateFail                             // 用户创建失败
	ErrorIncorrectUsernameOrPassword                // 用户名或密码错误
	ErrorGenerateTokenFail                          // 生成token失败
	ErrorUserIsDisabled                             // 用户已被封禁
	ErrorUserDeleteFail                             // 用户删除失败
	ErrorAiModelIsExist                             // AI模型已经存在
	ErrorAiModelMapCreateFail                       // AI模型关系映射创建失败
	ErrorAiModelMapUpdateFail                       // AI模型关系映射更新失败
	ErrorAiModelMapSelectFail                       // AI模型关系映射查询失败
	ErrorAiModelMapDeleteFail                       // AI模型关系映射删除失败
)
