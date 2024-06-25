package state

type States struct {
	Buttons        []ButtonState
	Sticks         []StickState
	UpdateCallback func()
}

func NewStates(buttons []ButtonState, sticks []StickState) States {
	return States{
		Buttons:        buttons,
		Sticks:         sticks,
		UpdateCallback: func() {},
	}
}
