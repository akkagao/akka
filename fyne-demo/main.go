package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Hello")

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("file",
			fyne.NewMenuItem("open           CM+x", func() {

			}),
			fyne.NewMenuItem("save", func() {}),
		),
		fyne.NewMenu("edit",
			fyne.NewMenuItem("del", func() {}),
			fyne.NewMenuItem("rename", func() {}),
		),
	))
	w.SetContent(widget.NewVBox(

		widget.NewLabel("Hello Fyne! -------- hahahh -------"),
		widget.NewButton("Quitsdf", func() {
			app.Quit()
		}),
		&widget.Button{Text: "hehe0", Style: 0},
		&widget.Button{Text: "hehe1", Style: 1},
		&widget.Button{Text: "hehe2", Style: 2},

		widget.NewProgressBar(),
		widget.NewRadio([]string{"Radio"}, func(args string) {
			fmt.Println(args)
		}),
		widget.NewHBox(widget.NewLabel("aa"), widget.NewLabel("bb")),
		widget.NewVBox(widget.NewLabel("aa1"), widget.NewLabel("bbb")),
	))

	w.ShowAndRun()
}
