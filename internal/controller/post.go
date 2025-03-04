package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"ztalk/internal/models"
	"ztalk/internal/service"
	"ztalk/pkg/request"
	"ztalk/pkg/response"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		response.Error(c, response.CodeInvalidParam)
		return
	}
	userID, err := request.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("request.GetCurrentUserID() failed", zap.Error(err))
		response.Error(c, response.CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt() failed", zap.Error(err))
		response.Error(c, response.CodeInvalidParam)
		return
	}
	data, err := service.GetPostDetail(id)
	if err != nil {
		zap.L().Error("service.GetPostDetail() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, data)
}

func GetPostListHandler(c *gin.Context) {
	page, size := request.GetPageInfo(c)
	data, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("service.GetPostList() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, data)
}
