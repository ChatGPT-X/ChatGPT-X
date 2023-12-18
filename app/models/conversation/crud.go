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

// Get 根据 ID 获取对话记录。
func Get(id string) (Conversation, error) {
	var conversation Conversation
	if err := model.DB.Where("id = ?", id).First(&conversation).Error; err != nil {
		return Conversation{}, err
	}
	return conversation, nil
}
