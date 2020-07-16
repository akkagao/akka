package testleak_demo

import (
	"fmt"
)

func demo2(cout int) {
	fmt.Println("demo")
	for i := 0; i < cout; i++ {
		fmt.Println(i)
	}
}