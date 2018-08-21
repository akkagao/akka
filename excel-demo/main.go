package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func mainTest() {
	xlsx := excelize.NewFile()

	for rowNow := 1; rowNow < 11; rowNow++ {
		fmt.Println("A" + strconv.Itoa(rowNow))
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(rowNow), "Hello A"+strconv.Itoa(rowNow))
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(rowNow), "Hello B"+strconv.Itoa(rowNow))
	}

	err := xlsx.SaveAs("./book1.xlsx")

	if err != nil {
		fmt.Println(err)
	}
}
