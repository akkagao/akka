package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

var testError = xerrors.New("test")
var testError2 = xerrors.New("test1")

func main() {
	fmt.Println(123)
	err := getError()
	if xerrors.Is(err, testError) {
		fmt.Println("true")
	}
}

func getError() error {
	return testError
}
