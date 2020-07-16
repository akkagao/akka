package main

import (
	"fmt"

	"github.com/akka/easyjson-demo/model"
)

//go:generate easyjson -all model/user.go
func main() {
	user := model.User{
		Id:      4,
		Name:    "haha",
		Address: 1,
	}
	b, _ := user.MarshalJSON()
	fmt.Println(string(b))
	s := `{"Id":4,"Name":"haha"}`
	u1 := model.User{}
	u1.UnmarshalJSON([]byte(s))

	b2, _ := u1.MarshalJSON()
	fmt.Println(string(b2))
}
