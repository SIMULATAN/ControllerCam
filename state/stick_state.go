package state

type StickState struct {
	Id   uint
	Name string
	// position from -1 to 1
	State float64
}
