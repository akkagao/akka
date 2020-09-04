package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/smallnest/rpcx/client"

	"prcx-demo/serviceInterface/user"
)

var userServiceClient client.XClient

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	userServiceClient = client.NewXClient("UserService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer userServiceClient.Close()

	// for i := 0; i < 1; i++ {
	start := time.Now()

	userRegisterReq := &user.RegisterUserReq{
		Name:     RandString(10),
		Barthday: "2020-09-15",
		Sex:      "man",
	}
	userRegisterRes := RegisterUser(userRegisterReq)

	fmt.Println(userRegisterRes.Id)

	// getUserReq := &user.GetUserReq{Id: userRegisterRes.Id}
	// getUserRes := &user.GetUserRes{}
	// userServiceClient.Call(context.Background(), "GetUser", getUserReq, getUserRes)

	// fmt.Printf("getUserRes:%v", getUserRes)

	fmt.Println(time.Since(start))

	// }
}

func RegisterUser(userRegisterReq *user.RegisterUserReq) *user.RegisterUserRes {
	userRegisterRes := &user.RegisterUserRes{}
	userServiceClient.Call(context.Background(), "RegisterUser", userRegisterReq, userRegisterRes)
	return userRegisterRes
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
