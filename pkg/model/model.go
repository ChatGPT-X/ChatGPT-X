package model

import (
	"chatgpt_x/pkg/config"
	"chatgpt_x/pkg/e"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB use Connect database(mysql).
func ConnectDB() *gorm.DB {
	var (
		err      error
		logLevel dbLogger.LogLevel
	)

	var (
		host     = config.GetString("databases.mysql.host")
		port     = config.GetString("databases.mysql.port")
		database = config.GetString("databases.mysql.database")
		username = config.GetString("databases.mysql.username")
		password = config.GetString("databases.mysql.password")
		charset  = config.GetString("databases.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	// determine the error level based on the debug
	logLevel = dbLogger.Silent
	if config.GetBool("app.debug") {
		logLevel = dbLogger.Info
	}

	// open a database connect
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: dbLogger.Default.LogMode(logLevel),
	})
	e.HasError(err)
	return DB
}
