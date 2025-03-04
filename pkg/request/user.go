package request

import (
	"github.com/gin-gonic/gin"
	"ztalk/pkg/message"
)

const CtxUserIDKey = "userID"

// GetCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
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
