package response

import (
	"ztalk/internal/models"
)

type PostListType struct {
	Code    ResCode              `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    []*models.PostDetail `json:"data"`    // 数据
}
