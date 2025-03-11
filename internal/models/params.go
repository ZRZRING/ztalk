package models

type SignUpParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VoteParam struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=-1 0 1"`
}

// PostListParam 获取帖子列表 query string 参数
type PostListParam struct {
	CommunityID int64  `json:"community_id" form:"community_id"`  // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`      // 页码
	Size        int64  `json:"size" form:"size" example:"10"`     // 每页数据量
	Order       string `json:"order" form:"order" example:"time"` // 排序依据
}

func (p PostListParam) GetStartEnd() (int64, int64) {
	return (p.Page - 1) * p.Size, p.Page * p.Size
}
