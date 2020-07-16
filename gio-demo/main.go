package main

import (
	// "gioui.org/ui/layout"
	// "gioui.org/ui/text"
	// "gioui.org/ui/measure"
	// "golang.org/x/image/font/gofont/goregular"
	// "golang.org/x/image/font/sfnt"
	// "gioui.org/ui"
    "gioui.org/ui/app"
)

func main() {
    go func() {
        w := app.NewWindow(nil)
        for range w.Events() {

        }
    }()
    app.Main()
}

// func main() {
//     go func() {
//         w := app.NewWindow(nil)
//         regular, _ := sfnt.Parse(goregular.TTF)
//         var cfg ui.Config
//         var faces measure.Faces
//         ops := new(ui.Ops)
//         for e := range w.Events() {
//             if e, ok := e.(app.DrawEvent); ok {
//                 cfg = &e.Config
//                 cs := layout.RigidConstraints(e.Size)
//                 ops.Reset()
//                 faces.Reset(cfg)

//                 // ADD YOUR LABELS
//                 lbl := text.Label{Face: faces.For(regular, ui.Sp(72)), Text: "Hello, World! 你哈卡"}
//                 lbl.Layout(ops, cs)

//                 w.Draw(ops)
//             }
//         }
//     }()
//     app.Main()
// }