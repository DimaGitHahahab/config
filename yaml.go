package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Timeout int    `yaml:"timeout"`
}

func GetConfig(yamlName string) (*Config, error) {
	data, err := os.ReadFile(yamlName)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
