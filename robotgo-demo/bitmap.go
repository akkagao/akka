package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	bitmap := robotgo.CaptureScreen(100, 220, 300, 40)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	// defer robotgo.FreeBitmap(bitmap)

	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------ ", fx, fy)

	// robotgo.SaveBitmap(bitmap, "test.png")
}
