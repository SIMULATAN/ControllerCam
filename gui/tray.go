package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func makeTray(window fyne.Window, a fyne.App, desk desktop.App) {
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
