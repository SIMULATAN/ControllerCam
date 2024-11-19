package gui

import (
	"controllercontrol/camera"
	"controllercontrol/state"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

func Cameras(_ fyne.Window, _ *state.States, handler *camera.ProtocolHandler) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return len((handler).GetCameras())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Camera")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			update(id, object, handler)
		},
	)

	go refresh(list)

	return list
}

func refresh(list *widget.List) {
	for {
		list.Refresh()
		time.Sleep(1 * time.Second)
	}
}

func update(id int, object fyne.CanvasObject, handler *camera.ProtocolHandler) {
	cam := handler.GetCameras()[id]
	var connectText string
	if cam.Conn.IsConnected() {
		connectText += "Connected!"
	} else {
		connectText += "Disconnected"
	}
	object.(*widget.Label).SetText(
		fmt.Sprintf("%v (%v): %v", cam.Config.Name, cam.Config.Host, connectText),
	)
}
