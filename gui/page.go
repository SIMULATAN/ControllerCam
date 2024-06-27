package gui

import (
	"controllercontrol/camera"
	cameras "controllercontrol/gui/pages/cameras"
	controllerdebug "controllercontrol/gui/pages/controller_debug"
	"controllercontrol/state"
	"fyne.io/fyne/v2"
)

type Page struct {
	Title      string
	View       func(w fyne.Window, states *state.States, handler *camera.ProtocolHandler) fyne.CanvasObject
	SupportWeb bool
}

var pages = map[string]Page{
	"controller_debug": {
		Title: "Controller Debug",
		View:  controllerdebug.ControllerDebug,
	},
	"cameras": {
		Title: "Cameras",
		View:  cameras.Cameras,
	},
}

var PageIndex = map[string][]string{
	"": {"controller_debug", "cameras"},
}
