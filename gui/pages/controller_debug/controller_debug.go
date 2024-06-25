package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ControllerDebug(w fyne.Window, states *state.States) fyne.CanvasObject {
	button, icons := buttonStates(states)
	stick, sliders := stickStates(states)
	states.UpdateCallback = func() {
		updateButtonStates(icons, states)
		updateStickStates(sliders, states)
	}
	tabs := container.NewAppTabs(
		container.NewTabItem("Buttons", button),
		container.NewTabItem("Sticks", stick),
	)
	return tabs
}
