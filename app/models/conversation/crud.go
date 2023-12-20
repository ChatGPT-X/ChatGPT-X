package conversation

import (
	"chatgpt_x/pkg/model"
	"errors"
	"gorm.io/gorm"
)

// Create 创建对话记录，通过 Conversation.ID 来判断是否创建成功。
func (m *Conversation) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新对话记录。
func (m *Conversation) Update() (rowsAffected int64, err error) {
	result := model.DB.Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// UpdateByConversationID 根据 conversationID 更新对话记录。
func (m *Conversation) UpdateByConversationID() (rowsAffected int64, err error) {
	conversationID := m.ConversationID
	m.ConversationID = ""
	result := model.DB.Where("conversation_id = ?", conversationID).Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// CreateOrUpdate 记录存在则更新（只更新 update_time 字段），记录不存在则创建。
func (m *Conversation) CreateOrUpdate() (err error) {
	_, err = Get(m.ConversationID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m.Create()
	}
	if err != nil {
		return err
	}
	conversationModel := Conversation{
		ConversationID: m.ConversationID,
		UpdateTime:     m.UpdateTime,
	}
	if _, err = conversationModel.UpdateByConversationID(); err != nil {
		return err
	}
	return nil
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

// HasConversationIDExist 通过 conversationID 判断对话记录是否存在，存在返回 true，不存在返回 false。
func HasConversationIDExist(conversationID string) bool {
	_, err := Get(conversationID)
	return err == nil
}

// List 获取对话记录列表。
func List(offset, limit int, userID uint, status string) ([]Conversation, int64, error) {
	var conversations []Conversation
	db := model.DB.
		Model(&Conversation{}).
		Where("user_id = ?", userID).
		Where("type = ?", TypeWeb).
		Where("status = ?", status).
		Omit("id", "user_id", "ai_token_id", "status").
		Order("update_time desc")
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Offset(offset).Limit(limit).Find(&conversations).Error; err != nil {
		return nil, 0, err
	}
	return conversations, count, nil
}
