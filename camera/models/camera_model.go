package camera_models

import (
	"controllercontrol/config"
	"controllercontrol/visca"
	"encoding/hex"
	"fmt"
	"strings"
)

type CameraModel interface {
	CreatePacket(payload []byte) []byte
	// Connect creates a NetworkConnection and connects asynchronously - connection errors will just be printed
	Connect(config config.CameraConfig) *visca.NetworkConnection

	RecallPreset(number uint8) []byte
	PanTilt(panSpeed, tiltSpeed, panPosition, tiltPosition int16, relative bool) []byte

	InterpretResponse(message []byte) string
}

const Terminator = 0xFF

// Encode Converts a signed integer to hex with each nibble seperated by a 0
func Encode(num int16) []byte {
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
