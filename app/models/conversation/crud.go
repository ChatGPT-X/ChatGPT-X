package conversation

import "chatgpt_x/pkg/model"

// Create 创建对话记录，通过 Conversation.ID 来判断是否创建成功。
func (m *Conversation) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新对话记录。
func (m *Conversation) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// CreateOrUpdate 记录存在则更新，记录不存在则创建。
func (m *Conversation) CreateOrUpdate() (err error) {
	return model.DB.Where(Conversation{ConversationID: m.ConversationID}).
		Assign(m).
		Omit("id").
		FirstOrCreate(&Conversation{}).Error
}

// Get 根据 conversationID 获取对话记录。
func Get(conversationID string) (Conversation, error) {
	var conversation Conversation
	if err := model.DB.
		Where("conversation_id = ?", conversationID).
		First(&conversation).Error; err != nil {
		return Conversation{}, err
	}
	return conversation, nil
}
