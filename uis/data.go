package uis

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DataScreen(_ fyne.Window) fyne.CanvasObject {
	counter := 0
	label := widget.NewLabel("Count is: 0")
	button := widget.NewButton("Counter", func() {
		counter += 1
		label.SetText(fmt.Sprintf("Count is: %d", counter))
	})

	content := container.NewVBox(label, button)

	return container.NewCenter(content)
}
