package model

import (
	"encoding/json"
	"fmt"
)

func UserToJson() string {
	user := User{
		Id:      10,
		Name:    "akka",
		Address: 1,
	}
	b, _ := json.Marshal(user)
	return string(b)
}

func UserToJsonEasy() string {
	user := User{
		Id:      10,
		Name:    "akka",
		Address: 1,
	}
	b, _ := user.MarshalJSON()
	fmt.Println(string(b))
	return string(b)
}

func JsonToUser(j string) User {
	user := User{}
	json.Unmarshal([]byte(j), &user)
	return user
}

func JsonToUserEasy(j string) User {
	user := User{}
	user.UnmarshalJSON([]byte(j))
	return user

}
