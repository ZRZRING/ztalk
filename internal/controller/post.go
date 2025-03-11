package controller

import (
	"strconv"
	"ztalk/internal/models"
	"ztalk/internal/service"
	"ztalk/pkg/request"
	"ztalk/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// GetPostsHandler 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.PostListParam false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.PostListType
// @Router /posts [get]
func GetPostsHandler(c *gin.Context) {
	p := &models.PostListParam{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler with invalid params", zap.Error(err))
		response.Error(c, response.CodeInvalidParam)
		return
	}
	data, err := service.GetPosts(p)
	if err != nil {
		zap.L().Error("service.GetPostList() failed", zap.Error(err))
		response.Error(c, response.CodeServerBusy)
		return
	}
	response.Success(c, data)
}
