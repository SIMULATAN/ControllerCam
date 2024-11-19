package camera_models

import (
	"controllercontrol/config"
	"controllercontrol/utils"
	"controllercontrol/visca"
)

var (
	SYS_Menu_ON          = []byte{0x01, 0x04, 0x06, 0x06, 0x02}
	CAM_Picture_Flip_ON  = []byte{0x01, 0x04, 0x66, 0x02}
	CAM_Picture_Flip_OFF = []byte{0x01, 0x04, 0x66, 0x03}

	RETURN_ACK                  = []byte{0x41}
	RETURN_Completion           = []byte{0x51}
	RETURN_SyntaxError          = []byte{0x60, 0x02}
	RETURN_CommandNotExecutable = []byte{0x61, 0x41}
)

func CAM_Brightness(p byte, q byte) []byte {
	return []byte{0x01, 0x04, 0xA1, 0x00, 0x00, p, q}
}

type RGBLink struct{}

func (s *RGBLink) Connect(config config.CameraConfig) *visca.NetworkConnection {
	return visca.Connect("tcp", config.Host)
}

func (s *RGBLink) CreatePacket(payload []byte) []byte {
	// Create the header
	source := 0x00
	destination := 0x01

	leftNibble := uint8(8) + uint8(source)
	rightNibble := uint8(destination)
	header := leftNibble<<4 + rightNibble

	// Construct the packet
	bytes := make([]byte, 0)
	bytes = append(bytes, header)
	bytes = append(bytes, payload...)
	bytes = append(bytes, Terminator)
	return bytes
}

func (s *RGBLink) RecallPreset(number uint8) []byte {
	return []byte{0x01, 0x04, 0x3F, 0x02, number}
}

func (s *RGBLink) InterpretResponse(bytes []byte) string {
	message := []byte{bytes[1]}
	if utils.BytesEqual(message, RETURN_ACK) {
		return "[SUC] ACK"
	} else if utils.BytesEqual(message, RETURN_Completion) {
		return "[SUC] Completion"
	} else if utils.BytesEqual(message, RETURN_SyntaxError) {
		return "[ERR] Syntax Error"
	} else if utils.BytesEqual(message, RETURN_CommandNotExecutable) {
		return "[ERR] Command not executable"
	}

	return utils.DumpByteSlice(message)
}
