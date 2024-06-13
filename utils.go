package main

func GetButtonState(state uint32, button uint32) bool {
	return (state>>button)&1 > 0
}
