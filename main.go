package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"

	"os"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Egg timer"),
			app.Size(unit.Dp(400), unit.Dp(600)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	// create new window

	// ops are the operations from the UI
	var ops op.Ops

	// startButton is a clickable widget
	var startButton widget.Clickable

	// th defnes the material design style
	th := material.NewTheme(gofont.Collection())
	// listen for events in the window.
	for e := range w.Events() {

		// detect what type of event
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		// this is sent when the application should re-render.
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			btn := material.Button(th, &startButton, "Start")
			btn.Layout(gtx)
			e.Frame(gtx.Ops)

		}
	}
	return nil
}
