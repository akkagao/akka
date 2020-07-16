package model

import (
	"fmt"
	"testing"
	"time"
)

func TestUserToJson(t *testing.T) {
	now := time.Now()
	for i := 0; i < 100000; i++ {
		UserToJson()
	}
	fmt.Println(time.Since(now))
}

func TestUserToJsonEasy(t *testing.T) {
	now := time.Now()
	for i := 0; i < 100000; i++ {
		UserToJsonEasy()
	}
	fmt.Println(time.Since(now))
}

func BenchmarkTestUserToJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserToJson()
	}
}

func BenchmarkUserToJsonEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserToJsonEasy()
	}
}

func TestJsonToUser(t *testing.T) {
	now := time.Now()
	for i := 0; i < 100000; i++ {
		JsonToUser(`{"Id":10,"Name":"akka","Address":1}`)
	}
	fmt.Println(time.Since(now))
}

func TestJsonToUserEasy(t *testing.T) {
	now := time.Now()
	for i := 0; i < 100000; i++ {
		JsonToUserEasy(`{"Id":10,"Name":"akka","Address":1}`)
	}
	fmt.Println(time.Since(now))
}
