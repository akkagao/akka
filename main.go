package main

import (
	"fmt"
)

func main() {
	// slice 如果没有扩容，slice的修改会修改底层数组
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	slice = append(slice, 6)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println(slice)
	fmt.Println(arr)
}
