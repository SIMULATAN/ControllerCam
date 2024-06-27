package gui

import (
	"controllercontrol/camera"
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ControllerDebug(_ fyne.Window, states *state.States, _ *camera.ProtocolHandler) fyne.CanvasObject {
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
