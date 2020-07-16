package main

import (
	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
)

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 768

	win := gi.NewMainWindow("gogi-tabview-test", "GoGi TabView Test", width, height)

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()

	tv := gi.AddNewTabView(mfr, "tv")
	tv.NewTabButton = true

	lbl1 := tv.AddNewTab(gi.KiT_Label, "This is Label1").(*gi.Label)
	lbl1.SetText("this is the contents of the first tab")
	lbl1.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

	lbl2 := tv.AddNewTab(gi.KiT_Label, "And this Label2").(*gi.Label)
	lbl2.SetText("this is the contents of the second tab")
	lbl2.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

	tv.SelectTabIndex(0)

	// main menu
	appnm := gi.AppName()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.ChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.SetCloseCleanFunc(func(w *gi.Window) {
		go gi.Quit() // once main window is closed, quit
	})

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()
}