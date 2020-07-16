package main

import (
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func main() {

	var err error

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Import example-pdf.pdf with gofpdi free pdf document importer
	tpl1 := gofpdi.ImportPage(pdf, "example-pdf.pdf", 1, "/MediaBox")

	pdf.AddPage()

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 50, 150, 0)

	pdf.SetFont("Helvetica", "", 20)
	pdf.Cell(0, 0, "Import existing PDF into gofpdf document with gofpdi")

	err = pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		panic(err)
	}
}
