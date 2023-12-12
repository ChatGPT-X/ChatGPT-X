package user_service

import (
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/app/models/user"
	"chatgpt_x/pkg/auth"
	"chatgpt_x/pkg/e"
	"fmt"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"time"
)

// UserService 用户服务。
type UserService struct{}

// DoRegister 用户注册。
func (s *UserService) DoRegister(paramsModel user.User) e.ErrInfo {
	// 检查该用户是否已经存在
	if user.HasByUsernameExist(paramsModel.Username) {
		return e.ErrInfo{
			Code: e.ErrorUserIsExist,
			Msg:  fmt.Errorf("username: %s is exist", paramsModel.Username),
		}
	}
	// 创建用户
	if err := paramsModel.Create(); err != nil {
		return e.ErrInfo{
			Code: e.ErrorUserCreateFail,
			Msg:  err,
		}
	}
	return e.ErrInfo{Code: e.SUCCESS}
}

// DoLogin 用户登录。
func (s *UserService) DoLogin(paramsModel user.User) (string, e.ErrInfo) {
	// 检查用户名和密码是否正确
	userModel, err := user.GetByUsername(paramsModel.Username)
	if err != nil || !user.CheckPassword(paramsModel.Password, userModel.Password) {
		return "", e.ErrInfo{
			Code: e.ErrorIncorrectUsernameOrPassword,
			Msg:  err,
		}
	}
	// 检查用户是否被封禁
	if userModel.Status == user.StatusDisable {
		return "", e.ErrInfo{
			Code: e.ErrorUserIsDisabled,
			Msg:  fmt.Errorf("user: %s is disabled, login fail", paramsModel.Username),
		}
	}
	// 获取用户 ai_token 信息，后面生成 jwt 的时候要记录该账号的 AI 密钥的类型
	aiTokenModel := ai_token.AiToken{}
	if userModel.AiTokenID != nil {
		aiTokenModel, _ = ai_token.Get(userModel.ID)
	}
	// 更新用户登录时间
	userModel.LastLoginTime = time.Now()
	if _, err = userModel.Update(); err != nil {
		return "", e.ErrInfo{
			Code: e.ErrorUserLoginFail,
			Msg:  err,
		}
	}
	// 生成 Token 授权
	token, err := auth.GenerateToken(auth.CustomClaims{
		UserID:   userModel.ID,
		UrlType:  aiTokenModel.Type,
		IsAdmin:  userModel.IsAdmin,
		Username: userModel.Username,
		Email:    userModel.Email,
	})
	if err != nil {
		return "", e.ErrInfo{
			Code: e.ErrorGenerateTokenFail,
			Msg:  err,
		}
	}
	return token, e.ErrInfo{Code: e.SUCCESS}
}

// List 查询用户列表。
func (s *UserService) List(page, pageSize int64) (paginator.Page[user.User], e.ErrInfo) {
	userModel := user.User{}
	pageData, err := userModel.List(page, pageSize)
	if err != nil {
		return paginator.Page[user.User]{}, e.ErrInfo{
			Code: e.ErrorUserSelectListFail,
			Msg:  err,
		}
	}
	data := pageData.(paginator.Page[user.User])
	return data, e.ErrInfo{Code: e.SUCCESS}
}

// Update 更新用户。
func (s *UserService) Update(paramsModel user.User) (int64, e.ErrInfo) {
	userModel, err := user.Get(paramsModel.ID)
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorUserSelectDetailFail,
			Msg:  err,
		}
	}
	userModel.AiTokenID = paramsModel.AiTokenID
	userModel.Password = paramsModel.Password
	userModel.Status = paramsModel.Status
	rows, err := userModel.Update()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorUserUpdateFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}

// Delete 删除用户。
func (s *UserService) Delete(id uint) (int64, e.ErrInfo) {
	userModel := user.User{
		ID: id,
	}
	rows, err := userModel.Delete()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorUserDeleteFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}
