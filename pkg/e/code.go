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
	ErrorUserLoginFail                              // 用户登录失败
	ErrorGenerateTokenFail                          // 生成token失败
	ErrorUserIsDisabled                             // 用户已被封禁
	ErrorUserDeleteFail                             // 用户删除失败
	ErrorAiModelIsExist                             // AI模型已经存在
	ErrorAiModelCreateFail                          // AI模型创建失败
	ErrorAiModelUpdateFail                          // AI模型更新失败
	ErrorAiModelSelectListFail                      // AI模型列表查询失败
	ErrorAiModelDeleteFail                          // AI模型删除失败
)
