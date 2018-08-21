package main

import "fmt"

func main() {
	fmt.Println("23")
	array := []int{1, 2, 465, 5, 78, 4}
	quickSort(array)
}

func quickSort(array []int) {
	fmt.Println(array)
	if len(array) < 2 {
		return
	}

}
