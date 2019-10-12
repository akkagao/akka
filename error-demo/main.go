package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("1")
	err2 := fmt.Errorf("2123 [%w]", err)
	fmt.Println(err2)
	fmt.Println(errors.Is(err2, err))
	err3 := errors.Unwrap(err2)
	if err == err3 {
		fmt.Println("success")
	}
	if errors.Is(err3, err) {
		fmt.Println("success")
	}

	if errors.Is(err2, err) {
		fmt.Println("is")
	}

	if !errors.Is(err, err2) {
		fmt.Println("! is")
	}

}
