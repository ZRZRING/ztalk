package request

import (
	"ztalk/internal/models"
	"ztalk/pkg/message"

	"github.com/gin-gonic/gin"
)

// GetCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(models.ContextUserIDKey)
	if !ok {
		err = message.ErrUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = message.ErrUserNotLogin
		return
	}
	return
}
