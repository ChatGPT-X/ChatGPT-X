package e

const (
	SUCCESS       = 200 // 成功
	ERROR         = 500 // 系统内部未知错误
	InvalidParams = 400 // 请求参数错误
	ErrorAuthFail = 401 //	认证失败或权限不足
)
const (
	ErrorGenerateTokenFail           = 10001 + iota // 生成token失败
	ErrorUserIsExist                                // 用户已经存在
	ErrorUserCreateFail                             // 用户创建失败
	ErrorIncorrectUsernameOrPassword                // 用户名或密码错误
	ErrorUserLoginFail                              // 用户登录失败
	ErrorUserIsDisabled                             // 用户已被封禁
	ErrorUserSelectListFail                         // 用户列表查询失败
	ErrorUserDeleteFail                             // 用户删除失败
	ErrorAiModelIsExist                             // AI模型已经存在
	ErrorAiModelCreateFail                          // AI模型创建失败
	ErrorAiModelUpdateFail                          // AI模型更新失败
	ErrorAiModelSelectListFail                      // AI模型列表查询失败
	ErrorAiModelDeleteFail                          // AI模型删除失败
	ErrorAiTokenIsExist                             // AI密钥已经存在
	ErrorAiTokenCreateFail                          // AI密钥创建失败
	ErrorAiTokenUpdateFail                          // AI密钥更新失败
	ErrorAiTokenSelectListFail                      // AI密钥列表查询失败
	ErrorAiTokenDeleteFail                          // AI密钥删除失败
)
