package bootstrap

import (
	"chatgpt_x/app/http/middlewares"
	"chatgpt_x/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRoute used for init Router and middleware.
func SetupRoute() http.Handler {
	engine := gin.New()
	//engine.Use(gin.Logger(), gin.Recovery())
	middlewares.Register(engine)
	routes.Register(engine)
	return engine
}
