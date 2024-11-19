package camera_models

import (
	"controllercontrol/config"
	"controllercontrol/utils"
	"controllercontrol/visca"
	"slices"
)

//https://marshall-usa.com/pdf/CV630-IP-RS-232_command.pdf

type Marshall struct{}

func (s *Marshall) Connect(config config.CameraConfig) *visca.NetworkConnection {
	return visca.Connect("udp", config.Host)
}

var sequenceNumber = 0

var terminator = []byte{0xFF}

func (s *Marshall) CreatePacket(payload []byte) []byte {
	// 81 = 8x | x = camera address
	payload = slices.Concat([]byte{0x81}, payload, terminator)
	payloadLength := []byte{byte(len(payload) >> 8), byte(len(payload) & 0xFF)}
	sequenceNumber++
	// convert sequence number to 4 bytes
	sequence := make([]byte, 4)
	sequence[0] = byte(sequenceNumber >> 24)
	sequence[1] = byte(sequenceNumber >> 16)
	sequence[2] = byte(sequenceNumber >> 8)
	sequence[3] = byte(sequenceNumber)
	return slices.Concat([]byte{0x01, 0x00}, payloadLength, sequence, payload)
}

func (s *Marshall) RecallPreset(number uint8) []byte {
	return []byte{0x01, 0x04, 0x3F, 0x02, number}
}

func (s *Marshall) PanTilt(panSpeed, tiltSpeed, panPosition, tiltPosition int16, relative bool) []byte {
	// Pan-tiltPosInq
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
	bytes = append(bytes, Encode(panPosition)...)
	bytes = append(bytes, Encode(tiltPosition)...)
	return bytes
}

func (s *Marshall) InterpretResponse(message []byte) string {
	return utils.DumpByteSlice(message)
}
