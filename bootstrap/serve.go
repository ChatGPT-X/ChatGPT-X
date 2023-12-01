package bootstrap

import (
	"chatgpt_x/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// SetupServe is used to initialize the http server
func SetupServe(r http.Handler) *http.Server {
	addr := config.GetString("http.listen_host") + ":" + config.GetString("http.listen_port")
	server := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       time.Duration(config.GetInt("http.read_timeout")) * time.Second,
		ReadHeaderTimeout: time.Duration(config.GetInt("http.read_header_timeout")) * time.Second,
		WriteTimeout:      time.Duration(config.GetInt("http.write_timeout")) * time.Second,
		IdleTimeout:       time.Duration(config.GetInt("http.idle_timeout")) * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Printf("[info] start http server listening %s", addr)
	return server
}

// getMode 根据配置返回一个对应的 Mode.
func getMode() string {
	if config.GetBool("app.debug") {
		return gin.DebugMode
	}
	return gin.ReleaseMode
}

// NewServe 获取一个新的服务.
func NewServe() *http.Server {
	return SetupServe(SetupRoute())
}
