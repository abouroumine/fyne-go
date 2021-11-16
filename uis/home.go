package uis

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Home Screen"),
		),
	)
}
