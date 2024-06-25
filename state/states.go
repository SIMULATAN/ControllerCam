package state

type States struct {
	Buttons        []ButtonState
	UpdateCallback func()
}

func NewStates(buttons []ButtonState) States {
	return States{
		Buttons:        buttons,
		UpdateCallback: func() {},
	}
}
