package controller

import (
	"strconv"
	"ztalk/internal/service"
	"ztalk/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	data, err := service.GetCommunityByID(id)
	if err != nil {
		zap.L().Error("service.GetCommunityByID() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, data)
}
