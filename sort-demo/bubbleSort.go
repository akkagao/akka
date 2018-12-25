package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := rand.Perm(10)
	// array := []int{1, 3, 4}
	fmt.Println("排序前：", array)
	bubbleSort(array)
	fmt.Println("排序后：", array)
}
func bubbleSort(array []int) {
	flag := false
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
