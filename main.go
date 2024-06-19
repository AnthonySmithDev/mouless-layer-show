package main

import (
	"flag"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	text := flag.String("mode", "", "The mode to display")
	flag.Parse()

	a := app.New()
	w := a.NewWindow("Mouseless GUI")

	w.SetContent(
		container.New(
			layout.NewCenterLayout(),
			widget.NewLabel(*text),
		),
	)

	go func() {
		time.Sleep(time.Second / 2)
		w.Close()
	}()

	w.ShowAndRun()
}
