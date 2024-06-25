package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func buttonStates(states *state.States) (fyne.CanvasObject, []*widget.Icon) {
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
			buttonState := states.Buttons[i]
			box := obj.(*fyne.Container)
			iconWidget := box.Objects[0].(*widget.Icon)
			icons[i] = iconWidget
			setStateIcon(iconWidget, buttonState)
			box.Objects[1].(*widget.Label).SetText(buttonState.Name)
		},
	)
	return container.NewStack(list), icons
}

func updateButtonStates(icons []*widget.Icon, states *state.States) {
	for i := range states.Buttons {
		setStateIcon(icons[i], states.Buttons[i])
	}
}

func setStateIcon(icon *widget.Icon, button state.ButtonState) {
	if button.State {
		icon.Resource = theme.ConfirmIcon()
	} else {
		icon.Resource = theme.CancelIcon()
	}
	icon.Refresh()
}
