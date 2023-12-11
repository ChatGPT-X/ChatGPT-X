package openai_service

type BackendApiService struct{}

//func (s *BackendApiService) Conversation(userID uint) e.ErrInfo {
//	type Result struct {
//		Token string
//	}
//systemSetting, err := system_setting.GetDetail()
//if err != nil {
//	return e.ErrInfo{
//		Code: e.ErrorSystemSettingSelectDetailFail,
//		Msg:  fmt.Errorf("create conversation fail: %s", err),
//	}
//}
//result := Result{}
//model.DB.Table("t_users AS u").
//	Joins("LEFT JOIN t_ai_tokens AS at ON at.id = u.ai_token_id").
//	Where("u.id = ?", userID).
//	Where("at.type = ?", ai_token.TypeAccessToken).
//	Select("at.token as token").
//	Scan(&result)
//fmt.Println(result)
//return e.ErrInfo{
//	Code: e.SUCCESS,
//	Msg:  nil,
//}
//}
