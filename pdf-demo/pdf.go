package main

import (
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.Image("aaa.jpg", 0, 0, nil) // print image

	// pdf.SetX(250) // move current location
	// pdf.SetY(200)

	pdf.WritePdf("image.pdf")
}
