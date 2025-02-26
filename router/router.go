package router

import (
	"net/http"
	"ztalk/internal/controller"
	"ztalk/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	{
		v1 := r.Group("/api/v1")
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 page not found")
	})
	return r
}
