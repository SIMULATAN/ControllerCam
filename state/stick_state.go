package state

type StickState struct {
	Id   int
	Name string
	// position from -1 to 1
	State float64
}
