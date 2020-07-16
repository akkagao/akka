package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {

	for {
		mleft := robotgo.AddEvent("mleft")
		if mleft == 0 {
			x, y := robotgo.GetMousePos()
			fmt.Println("pos: ", x, y)
		}
	}

	// fmt.Println("==============================")
	// robotgo.GetScreenSize()
	// bitmap := robotgo.CaptureScreen(1092, 623, 35, 35)
	// // bitmap := robotgo.CaptureScreen(2184, 1246, 305, 305)

	// // defer robotgo.FreeBitmap(bitmap)

	// fmt.Println("...", bitmap)

	// fx, fy := robotgo.FindBitmap(bitmap)
	// fmt.Println("FindBitmap------", fx, fy)

	// robotgo.SaveBitmap(bitmap, "test.png")

	// time.Sleep(time.Second * 5)
	// fmt.Println("=========")
	// img := robotgo.OpenBitmap("./oneyuan.png")
	// x, y := robotgo.FindBitmap(img)
	// fmt.Println(x, y)
}

// pos:  1098 622
// pos:  1098 622
// pos:  1095 649
// pos:  1095 649
// pos:  1126 627
// pos:  1126 627
// pos:  1130 647
// pos:  1130 647
