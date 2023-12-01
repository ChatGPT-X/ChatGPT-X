package e

const (
	SUCCESS                          = 200   // 成功
	ERROR                            = 500   // 错误
	InvalidParams                    = 400   // 参数错误
	ErrorUserIsExist                 = 10001 // 用户已经存在
	ErrorUserCreateFail              = 10002 // 用户创建失败
	ErrorIncorrectUsernameOrPassword = 10003 // 用户名或密码错误
	ErrorUserIsDisabled              = 10004 // 用户已被禁用
)
