package config

type MappingsConfig struct {
	Presets []PresetMapping
}

type PresetMapping struct {
	Preset uint8
	Button string
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
