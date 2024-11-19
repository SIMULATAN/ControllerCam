package camera_models

import (
	"controllercontrol/config"
	"controllercontrol/visca"
)

type CameraModel interface {
	CreatePacket(payload []byte) []byte
	// Connect creates a NetworkConnection and connects asynchronously - connection errors will just be printed
	Connect(config config.CameraConfig) *visca.NetworkConnection

	RecallPreset(number uint8) []byte

	InterpretResponse(message []byte) string
}

const Terminator = 0xFF
