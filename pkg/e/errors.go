package e

import "chatgpt_x/pkg/logger"

var MsgFlags = map[int]string{
	SUCCESS:       "success",
	ERROR:         "系统内部未知错误",
	InvalidParams: "请求参数错误",
	ErrorAuthFail: "认证失败或权限不足",

	ErrorGenerateTokenFail:           "生成token失败",
	ErrorUserIsExist:                 "用户已经存在",
	ErrorUserCreateFail:              "用户创建失败",
	ErrorIncorrectUsernameOrPassword: "用户名或密码错误",
	ErrorUserLoginFail:               "用户登录失败",
	ErrorUserIsDisabled:              "用户已被禁用",
	ErrorUserSelectDetailFail:        "用户详情获取失败",
	ErrorUserSelectListFail:          "用户列表获取失败",
	ErrorUserUpdateFail:              "用户更新失败",
	ErrorUserDeleteFail:              "用户删除失败",
	ErrorAiModelIsExist:              "AI模型已经存在",
	ErrorAiModelCreateFail:           "AI模型创建失败",
	ErrorAiModelUpdateFail:           "AI模型更新失败",
	ErrorAiModelSelectDetailFail:     "AI模型详情获取失败",
	ErrorAiModelSelectListFail:       "AI模型列表获取失败",
	ErrorAiModelDeleteFail:           "AI模型删除失败",
	ErrorAiTokenIsDisabled:           "AI密钥已被禁用",
	ErrorAiTokenIsExist:              "AI密钥已经存在",
	ErrorAiTokenCreateFail:           "AI密钥创建失败",
	ErrorAiTokenUpdateFail:           "AI密钥更新失败",
	ErrorAiTokenSelectDetailFail:     "AI密钥详情获取失败",
	ErrorAiTokenSelectListFail:       "AI密钥列表获取失败",
	ErrorAiTokenDeleteFail:           "AI密钥删除失败",
	ErrorSettingUpdateFail:           "系统设置更新失败",
	ErrorSettingSelectDetailFail:     "系统设置详情获取失败",

	// -------------- start --------------
	ErrorOpenaiInvalidParams: "Invalid Params, please contact the developer.",
	ErrorOpenaiRequestFail:   "Request for openai failed, please wait and try again.",
	// --------------  end  --------------
}

// ErrInfo 错误信息结构。
type ErrInfo struct {
	Code int
	Msg  error
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
		logger.Error(err)
		return true
	}
	return false
}

// CheckErrorFromCode 检查错误码是否为成功。
func CheckErrorFromCode(errCode int) bool {
	return errCode == SUCCESS
}
