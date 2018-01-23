package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
}

func NewConfigFromToml(filepath string) (*Config, error) {
	config := &Config{}
	if _, err := toml.Decode(filepath, config); err != nil {
		return nil, err
	}
	return config, nil
}
