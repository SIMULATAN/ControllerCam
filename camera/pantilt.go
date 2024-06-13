package camera

import (
	"controllercontrol/utils"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	CamMax = 24
)

func PanTilt(panSpeed, tiltSpeed, panPosition, tiltPosition int16, relative bool) []byte {
	if utils.Abs(panSpeed) > CamMax {
		panSpeed = CamMax * panSpeed / utils.Abs(panSpeed)
	}
	if utils.Abs(tiltSpeed) > CamMax {
		tiltSpeed = CamMax * tiltSpeed / utils.Abs(tiltSpeed)
	}

	panSpeedByte := byte(utils.Abs(panSpeed))
	tiltSpeedByte := byte(utils.Abs(tiltSpeed))

	var relativeByte byte
	if relative {
		relativeByte = 0x03
	} else {
		relativeByte = 0x02
	}

	bytes := []byte{
		0x01,
		0x06,
		relativeByte,
		panSpeedByte,
		tiltSpeedByte,
	}
	bytes = append(bytes, encode(panPosition)...)
	bytes = append(bytes, encode(tiltPosition)...)
	return bytes
}

// Converts a signed integer to hex with each nibble seperated by a 0
func encode(num int16) []byte {
	// """Converts a signed integer to hex with each nibble seperated by a 0"""
	// 200 => c8 in hex => [00, 00, 0c, 08]
	// Convert to hex
	position := int32(num)
	if position < 0 {
		// converts signed to unsigned
		position += 65536
	}

	hexString := fmt.Sprintf("%x", position)
	result := strings.Repeat("00", 4-len(hexString))
	for _, str := range hexString {
		result += "0" + string(str)
	}
	bytes, _ := hex.DecodeString(result)
	return bytes
}
