package uis

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const preference = "FyneApp"

type AppSideNavBar struct {
	Title, Introduction string
	View                func(w fyne.Window) fyne.CanvasObject
}

var (
	Menus = map[string]AppSideNavBar{
		"Home":     {"Home", "Welcome To Home", HomeScreen},
		"Data":     {"Data", "Welcome To Data", DataScreen},
		"Settings": {"Settings", "Welcome To Settings", SettingsScreen},
	}

	MenuIndex = map[string][]string{
		"": {"Home", "Data", "Settings"},
	}
)

func SideNavBar(setMenu func(menu AppSideNavBar)) fyne.CanvasObject {
	a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid widget.TreeNodeID) []string {
			return MenuIndex[uid]
		},
		IsBranch: func(uid widget.TreeNodeID) bool {
			children, ok := MenuIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("")
		},
		UpdateNode: func(uid widget.TreeNodeID, branch bool, obj fyne.CanvasObject) {
			t, ok := Menus[uid]
			if !ok {
				fyne.LogError("Error "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Menus[uid]; ok {
				a.Preferences().SetString(preference, uid)
				setMenu(t)
			}
		},
	}
	return container.NewBorder(nil, nil, nil, nil, tree)
}
