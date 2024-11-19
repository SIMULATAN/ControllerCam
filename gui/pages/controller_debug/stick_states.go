package gui

import (
	"controllercontrol/state"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func stickStates(states *state.States) (fyne.CanvasObject, map[string]*widget.ProgressBar) {
	sliders := map[string]*widget.ProgressBar{}
	sliderIndexToName := make([]string, 0)
	for name := range states.Sticks {
		sliderIndexToName = append(sliderIndexToName, name)
	}
	widgets := make([]fyne.CanvasObject, 0)
	for i, stick := range states.Sticks {
		slider := widget.NewProgressBar()
		slider.Min = -1
		slider.Max = 1
		sliders[i] = slider
		label := widget.NewLabel(stick.Name)
		widgets = append(widgets, container.NewVBox(label, slider))
	}
	return container.NewVBox(widgets...), sliders
}

func updateStickStates(sliders map[string]*widget.ProgressBar, states *state.States) {
	for i, stick := range states.Sticks {
		sliders[i].SetValue(stick.State)
	}
}
