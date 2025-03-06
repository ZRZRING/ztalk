package controller

import (
	"ztalk/internal/models"
	"ztalk/internal/service"
	"ztalk/pkg/request"
	"ztalk/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(c *gin.Context) {
	p := new(models.VoteParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		checkValidator(c, err)
		return
	}
	userID, err := request.GetCurrentUserID(c)
	if err != nil {
		response.Error(c, response.CodeNeedLogin)
	}
	if err := service.Vote(userID, p); err != nil {
		zap.L().Error("service.Vote() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, nil)
}
