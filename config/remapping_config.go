package config

type Button struct {
	Id              int     `yaml:"id"`
	Name            *string `yaml:"name"`
	IsImplicitStick *bool   `yaml:"isImplicitStick"`
}

type Remappings map[string]Button
