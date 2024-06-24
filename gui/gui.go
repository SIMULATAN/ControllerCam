package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func RunGui() error {
	a := app.NewWithID("me.simulatan.controllercam")
	window := a.NewWindow("ControllerCam")
	path, err := fyne.LoadResourceFromPath("logo.png")
	if err != nil {
		return err
	}
	a.SetIcon(path)

	if desk, ok := a.(desktop.App); ok {
		menu := fyne.NewMenu("ControllerCam",
			fyne.NewMenuItem("Show", func() {
				window.Show()
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayMenu(menu)
	}

	window.SetContent(widget.NewLabel("Hello World!"))
	window.SetCloseIntercept(func() {
		window.Hide()
	})

	window.ShowAndRun()

	return nil
}
