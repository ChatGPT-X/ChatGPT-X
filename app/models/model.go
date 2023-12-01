package models

//// IdModel 基础 Model,后续的业务表自动加入 ID.
//type IdModel struct {
//	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
//}
//
//// DateModel 基础 Model,后续的业务表自动加入增改时间.
//type DateModel struct {
//	CreateAt time.Time `gorm:"column:created_at;index;autoCreateTime"`
//	UpdateAt time.Time `gorm:"column:updated_at;index;autoUpdateTime"`
//}

type CustomModel map[string]interface{}
