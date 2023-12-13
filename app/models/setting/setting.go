package setting

import "time"

// 设置表
type Setting struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	ApiBaseUrl string    `gorm:"column:api_base_url;type:varchar(255);NOT NULL" json:"api_base_url"`                      // api 基础地址
	ApiProxy   string    `gorm:"column:api_proxy;type:varchar(255);NOT NULL" json:"api_proxy"`                            // api 代理地址
	ApiTimeout uint      `gorm:"column:api_timeout;type:tinyint(4) unsigned;NOT NULL" json:"api_timeout"`                 // api 超时时间（单位：秒）
	WebBaseUrl string    `gorm:"column:web_base_url;type:varchar(255);NOT NULL" json:"web_base_url"`                      // web 基础地址
	WebProxy   string    `gorm:"column:web_proxy;type:varchar(255);NOT NULL" json:"web_proxy"`                            // web 代理地址
	WebTimeout uint      `gorm:"column:web_timeout;type:tinyint(4) unsigned;NOT NULL" json:"web_timeout"`                 // web 超时时间（单位：秒）
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *Setting) TableName() string {
	return "t_settings"
}
