package bootstrap

import (
	"chatgpt_x/pkg/config"
	"chatgpt_x/pkg/model"
	"time"
)

// SetupMySQL used for init databases and ORM
func SetupMySQL() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	// set max db connections number
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// set max db idle connections number
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// expiration time for each connection
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")))
}
