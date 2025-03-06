package message

import "errors"

var (
	ErrUserExist       = errors.New("用户已存在")
	ErrUserNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrInvalidID       = errors.New("无效的ID")
	ErrUserNotLogin    = errors.New("用户未登录")
	ErrVoteTimeExpire  = errors.New("投票时间已过")
	ErrVoteRepeated    = errors.New("不允许重复投票")
)
