package main

import "fmt"
import "rsc.io/pdf"

func main() {
	file, err := pdf.Open("example-pdf.pdf")
    if err != nil {
        panic(err)
    }
    fmt.Println(file.NumPage())
}