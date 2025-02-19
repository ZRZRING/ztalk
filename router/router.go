package router

import (
	"net/http"
	"ztalk/app/controller"
	"ztalk/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 page not found")
	})
	return r
}
