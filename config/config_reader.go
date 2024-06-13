package config

import (
	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

func Read() (*Config, error) {
	content, err := os.ReadFile("controllercam.yml")
	if err == nil {
		return readConfig(content)
	} else {
		log.Debug().Msg("No config file found in current directory")
	}

	configFilePath, err := xdg.SearchConfigFile("controllercam/config.yml")
	if err == nil {
		content, err = os.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}
		return readConfig(content)
	} else {
		log.Debug().Msg("No config file found in any XDG config directory")
	}

	log.Debug().Msg("No config files found anywhere. Using default. THIS PROBABLY WILL NOT WORK FOR YOU!")
	return NewConfig(), nil
}

func readConfig(content []byte) (*Config, error) {
	config := NewConfig()
	err := yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
