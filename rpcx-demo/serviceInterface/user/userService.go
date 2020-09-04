package user

import (
	"context"
)

type UserService interface {
	// 注册用户
	RegisterUser(c context.Context, req *RegisterUserReq, res *RegisterUserRes) error
	// 查询用户
	GetUser(c context.Context, req *GetUserReq, res *GetUserRes) error
}
