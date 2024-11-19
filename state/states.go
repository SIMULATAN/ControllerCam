package state

import (
	"errors"
)

type States struct {
	Buttons        map[string]*ButtonState
	Sticks         map[string]*StickState
	UpdateCallback func()
}

func NewStates(buttons []ButtonState, sticks []StickState) States {
	buttonMap := make(map[string]*ButtonState)
	for _, button := range buttons {
		buttonMap[button.Name] = &button
	}
	stickMap := make(map[string]*StickState)
	for _, stick := range sticks {
		stickMap[stick.Name] = &stick
	}
	return States{
		Buttons:        buttonMap,
		Sticks:         stickMap,
		UpdateCallback: func() {},
	}
}

var noMeasurement = errors.New("no measurement")

func (c *States) GetButtonState(button string) (bool, error) {
	if button := c.Buttons[button]; button != nil {
		return button.State, nil
	}
	return false, noMeasurement
}

func (c *States) GetStickState(button string) (float64, error) {
	if stick := c.Sticks[button]; stick != nil {
		return stick.State, nil
	}
	return 0, noMeasurement
}
