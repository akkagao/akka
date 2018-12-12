package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// fmt.Println(time.Now().Nanosecond())
	b := rand.Perm(20)
	fmt.Println(b)
}
