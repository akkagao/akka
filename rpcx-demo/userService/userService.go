package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/smallnest/rpcx/server"

	"prcx-demo/serviceInterface/user"
)

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUser(c context.Context, req *user.GetUserReq, res *user.GetUserRes) error {
	fmt.Printf("获取ID：%v 的用户", req.Id)
	res.Name = "CrazyWolf"
	res.Sex = "nan"
	res.Barthday = "2020-09-02"
	return nil
}

func (u *UserServiceImpl) RegisterUser(c context.Context, req *user.RegisterUserReq, res *user.RegisterUserRes) error {
	fmt.Println(req.Name, req.Barthday, req.Sex)
	res.Id = rand.Int63n(1000000)
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("UserService", new(UserServiceImpl), "")
	s.Serve("tcp", ":8972")
}
