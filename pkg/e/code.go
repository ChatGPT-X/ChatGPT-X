package e

const (
	SUCCESS       = 200 // 成功
	ERROR         = 500 // 系统内部未知错误
	InvalidParams = 400 // 请求参数错误
	ErrorAuthFail = 401 //	认证失败或权限不足
)
const (
	ErrorGenerateTokenFail             = 10001 + iota // 生成token失败
	ErrorUserIsExist                                  // 用户已经存在
	ErrorUserCreateFail                               // 用户创建失败
	ErrorIncorrectUsernameOrPassword                  // 用户名或密码错误
	ErrorUserLoginFail                                // 用户登录失败
	ErrorUserIsDisabled                               // 用户已被禁用
	ErrorUserSelectDetailFail                         // 用户详情查询失败
	ErrorUserSelectListFail                           // 用户列表查询失败
	ErrorUserUpdateFail                               // 用户更新失败
	ErrorUserDeleteFail                               // 用户删除失败
	ErrorAiModelIsExist                               // AI模型已经存在
	ErrorAiModelCreateFail                            // AI模型创建失败
	ErrorAiModelUpdateFail                            // AI模型更新失败
	ErrorAiModelSelectDetailFail                      // AI模型详情查询失败
	ErrorAiModelSelectListFail                        // AI模型列表查询失败
	ErrorAiModelDeleteFail                            // AI模型删除失败
	ErrorAiTokenIsDisabled                            // AI密钥已被禁用
	ErrorAiTokenIsExist                               // AI密钥已经存在
	ErrorAiTokenCreateFail                            // AI密钥创建失败
	ErrorAiTokenUpdateFail                            // AI密钥更新失败
	ErrorAiTokenSelectDetailFail                      // AI密钥详情查询失败
	ErrorAiTokenSelectListFail                        // AI密钥列表查询失败
	ErrorAiTokenDeleteFail                            // AI密钥删除失败
	ErrorSystemSettingUpdateFail                      // 系统设置更新失败
	ErrorSystemSettingSelectDetailFail                // 系统设置详情查询失败
	// --------- OpenAI Start ---------
	ErrorSendRequestFail             // 发送请求失败
	ErrorAcceptResponseFail          // 接收响应失败
	ErrorChangeConversationTitleFail // openai-修改对话标题失败
	// --------- OpenAI End ---------
)
