package camera

import (
	"controllercontrol/utils"
	"strconv"
)

// Zoom factor is from -7 to 7
func Zoom(factor int8) []byte {
	var computedByte string
	if factor == 0 {
		computedByte = "0"
	} else if factor > 0 {
		computedByte = "2"
	} else {
		computedByte = "3"
	}
	computedByte += strconv.Itoa(int(utils.Abs(int16(factor))))

	result := []byte{
		0x01,
		0x04,
		0x07,
		utils.ParseSingleByte(computedByte),
	}
	return result
}
