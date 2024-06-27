package gui

import (
	"controllercontrol/camera"
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

var topWindow fyne.Window

func RunGui(states *state.States, handler *camera.ProtocolHandler) error {
	a := app.NewWithID("me.simulatan.controllercam")
	window := a.NewWindow("ControllerCam")
	topWindow = window
	path, err := fyne.LoadResourceFromPath("logo.png")
	if err != nil {
		return err
	}
	a.SetIcon(path)

	if desk, ok := a.(desktop.App); ok {
		makeTray(window, a, desk)
	}

	content := container.NewStack()

	setContent := func(p Page) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(p.Title)
			window = child
			child.SetContent(p.View(topWindow, states, handler))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = window
			})
			return
		}

		content.Objects = []fyne.CanvasObject{p.View(window, states, handler)}
		content.Refresh()
	}

	split := container.NewHSplit(makeSidebar(handler, setContent), content)
	// give the nav 20% of the window width
	split.Offset = 0.2
	window.SetContent(split)
	window.SetCloseIntercept(func() {
		window.Hide()
	})

	window.ShowAndRun()

	return nil
}
