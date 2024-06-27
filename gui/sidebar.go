package gui

import (
	"controllercontrol/camera"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeSidebar(handler *camera.ProtocolHandler, setContent func(page Page)) fyne.CanvasObject {
	status := makeStatus(handler)
	nav := makeNav(setContent)
	return container.NewBorder(nil, status, nil, nil, nav)
}

func makeStatus(handler *camera.ProtocolHandler) fyne.CanvasObject {
	label := widget.NewLabel(getName(handler))
	handler.ActiveCamera.AddListener(binding.NewDataListener(func() {
		label.SetText(getName(handler))
	}))
	return label
}

func getName(handler *camera.ProtocolHandler) string {
	cam := handler.GetActiveCamera()
	var name string
	if cam == nil {
		name = "No camera selected"
	} else {
		name = cam.Config.Name
	}
	return name
}
