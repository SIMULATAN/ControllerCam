package utils

func Abs(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func GetButtonState(state uint32, button uint32) bool {
	return (state>>button)&1 > 0
}
