package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func buttonStates(states *state.States) fyne.CanvasObject {
	icons := make([]*widget.Icon, len(states.Buttons))
	list := widget.NewList(
		func() int {
			return len(states.Buttons)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.CancelIcon())
			label := widget.NewLabel("Button")
			box := container.NewHBox(icon, label)
			return box
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			button := states.Buttons[i]
			box := obj.(*fyne.Container)
			iconWidget := box.Objects[0].(*widget.Icon)
			icons[i] = iconWidget
			setStateIcon(icons, i, states)
			box.Objects[1].(*widget.Label).SetText(button.Name)
		},
	)
	states.UpdateCallback = func() {
		for i := range states.Buttons {
			setStateIcon(icons, i, states)
		}
	}
	return container.NewStack(list)
}

func setStateIcon(icons []*widget.Icon, i int, states *state.States) {
	icon := icons[i]
	button := states.Buttons[i]
	if button.State {
		icon.Resource = theme.ConfirmIcon()
	} else {
		icon.Resource = theme.CancelIcon()
	}
	icon.Refresh()
}
