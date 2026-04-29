package config

type Button struct {
	Id              int     `yaml:"id"`
	Name            *string `yaml:"name"`
	IsImplicitStick *bool   `yaml:"isImplicitStick"`
	Invert          *bool   `yaml:"invert"`
}

type Remappings map[string]Button
