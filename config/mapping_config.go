package config

type MappingsConfig struct {
	Presets []PresetMapping
	Cameras []CameraMapping
}

type Mapping interface {
	GetButton() string
}

type PresetMapping struct {
	Preset uint8
	Button string
}

func (m *PresetMapping) GetButton() string {
	return m.Button
}

type CameraMapping struct {
	Camera string
	Button string
}

func (m *CameraMapping) GetButton() string {
	return m.Button
}

type Input struct {
	Index uint8
	Type  uint8
	// Data arbitrary data to be used by the interpreter
	Data interface{}
}

const (
	InputType_Button        = 0
	InputType_Stick         = 1
	InputType_StickAsButton = 2
)
