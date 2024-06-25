package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func stickStates(states *state.States) (fyne.CanvasObject, map[int]*widget.ProgressBar) {
	sliders := map[int]*widget.ProgressBar{}
	widgets := make([]fyne.CanvasObject, len(states.Sticks))
	for i, stick := range states.Sticks {
		slider := widget.NewProgressBar()
		slider.Min = -1
		slider.Max = 1
		sliders[i] = slider
		label := widget.NewLabel(stick.Name)
		widgets[i] = container.NewVBox(label, slider)
	}
	return container.NewVBox(widgets...), sliders
}

func updateStickStates(sliders map[int]*widget.ProgressBar, states *state.States) {
	for i, stick := range states.Sticks {
		sliders[i].SetValue(stick.State)
	}
}
