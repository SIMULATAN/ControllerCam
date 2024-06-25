package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
)

func ControllerDebug(w fyne.Window, states *state.States) fyne.CanvasObject {
	return buttonStates(states)
}
