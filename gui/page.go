package gui

import (
	gui "controllercontrol/gui/pages/controller_debug"
	"controllercontrol/state"
	"fyne.io/fyne/v2"
)

type Page struct {
	Title      string
	View       func(w fyne.Window, states *state.States) fyne.CanvasObject
	SupportWeb bool
}

var pages = map[string]Page{
	"controller_debug": {
		Title: "Controller Debug",
		View:  gui.ControllerDebug,
	},
}

var PageIndex = map[string][]string{
	"": {"controller_debug"},
}
