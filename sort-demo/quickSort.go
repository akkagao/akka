package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := rand.Perm(20)
	fmt.Println("未排序", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println("排序后", array)
}
func quickSort(array []int, low, heigh int) {
	if low > heigh {
		return
	}
	i, j := low, heigh
	temp := array[low]
	for i < j {
		for temp <= array[j] && i < j {
			j--
		}
		for temp >= array[i] && i < j {
			i++
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i], array[low] = array[low], array[i]
	quickSort(array, low, j-1)
	quickSort(array, j+1, heigh)
}
