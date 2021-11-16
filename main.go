package main

import (
	"abouroumine/fyne-app/uis"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("fyne.app")
	w := a.NewWindow("")
	title := widget.NewLabel("")
	intro := widget.NewLabel("")
	intro.Wrapping = fyne.TextWrapWord
	content := container.NewMax()
	setApp := func(t uis.AppSideNavBar) {
		title.SetText(t.Title)
		intro.SetText(t.Introduction)
		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}
	application := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	split := container.NewHSplit(uis.SideNavBar(setApp), application)
	split.Offset = 0.2
	w.SetContent(split)
	w.Resize(fyne.NewSize(1440, 640))
	w.ShowAndRun()
}
