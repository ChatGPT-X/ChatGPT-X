package e

const (
	SUCCESS                          = 200   // 成功
	ERROR                            = 500   // 系统内部未知错误
	InvalidParams                    = 400   // 请求参数错误
	ErrorAuthFail                    = 401   //	认证失败或权限不足
	ErrorUserIsExist                 = 10001 // 用户已经存在
	ErrorUserCreateFail              = 10002 // 用户创建失败
	ErrorIncorrectUsernameOrPassword = 10003 // 用户名或密码错误
	ErrorUserIsDisabled              = 10004 // 用户已被封禁
	ErrorAiModelIsExist              = 10005 // AI模型已经存在
	ErrorAiModelMapCreateFail        = 10006 // AI模型关系映射创建失败
	ErrorAiModelMapUpdateFail        = 10007 // AI模型关系映射更新失败
	ErrorAiModelMapSelectFail        = 10008 // AI模型关系映射查询失败
)
