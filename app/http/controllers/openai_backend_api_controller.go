package controllers

// OpenaiBackendApiController OPENAI WEB 接口控制器。
type OpenaiBackendApiController struct {
	BaseController
}

//// Conversation WEB 平台对话。
//func (oba *OpenaiBackendApiController) Conversation(c *gin.Context) {
//	userID := getUserID(c)
//	openaiService := openai_service.BackendApiService{}
//	openaiService.Conversation(userID)
//	//apiUrl := systemSetting.WebBaseUrl + "/backend-api/conversation"
//	//fmt.Println(aiToken.Token, apiUrl)
//}
