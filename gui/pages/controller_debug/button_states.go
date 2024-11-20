package gui

import (
	"controllercontrol/state"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"slices"
)

func buttonStates(states *state.States) (fyne.CanvasObject, map[string]*widget.Icon) {
	iconIndexToName := make([]string, 0)
	for name, states := range states.Buttons {
		if states.Id == -1 {
			fmt.Println("ID of button", name, "is -1, ignoring")
			continue
		}
		iconIndexToName = append(iconIndexToName, name)
	}
	// go iteration order is not guaranteed, so we sort the names
	slices.Sort(iconIndexToName)
	icons := make(map[string]*widget.Icon)

	list := widget.NewList(
		func() int {
			return len(iconIndexToName)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.CancelIcon())
			label := widget.NewLabel("Button")
			box := container.NewHBox(icon, label)
			return box
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			name := iconIndexToName[i]
			buttonState := states.Buttons[name]
			box := obj.(*fyne.Container)
			iconWidget := box.Objects[0].(*widget.Icon)
			icons[name] = iconWidget
			setStateIcon(iconWidget, *buttonState)
			box.Objects[1].(*widget.Label).SetText(buttonState.Name)
		},
	)
	return container.NewStack(list), icons
}

func updateButtonStates(icons map[string]*widget.Icon, states *state.States) {
	for name, s := range states.Buttons {
		if icons[name] == nil {
			//some things may be unmapped
			continue
		}
		setStateIcon(icons[name], *s)
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
