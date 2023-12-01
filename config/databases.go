package config

import "chatgpt_x/pkg/config"

// init this is the db configuration.
func init() {
	config.Add("databases", config.StrMap{
		"mysql": map[string]interface{}{
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"username": config.Env("DB_USERNAME", "root"),
			"password": config.Env("DB_PASSWORD", "root"),
			"database": config.Env("DB_DATABASE", "db_chatgpt_x"),
			"charset":  "utf8mb4",
			// db parameter settings
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
		"redis": map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),
			"db":       config.Env("REDIS_DB", 0),
		},
	})
}
