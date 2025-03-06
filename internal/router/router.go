package router

import (
	"net/http"
	"ztalk/internal/controller"
	"ztalk/pkg/jwt"
	"ztalk/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) (r *gin.Engine) {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	{
		v1 := r.Group("/api/v1")
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		v1.Use(jwt.Auth())
		{
			v1.GET("/community", controller.CommunityHandler)
			v1.GET("/community/:id", controller.CommunityDetailHandler)
			v1.GET("/post/:id", controller.GetPostDetailHandler)
			v1.GET("/posts/", controller.GetPostListHandler)
			v1.POST("/post", controller.CreatePostHandler)
			v1.POST("/vote", controller.VoteHandler)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 page not found")
	})
	return
}
