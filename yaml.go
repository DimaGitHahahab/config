package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port    int64  `yaml:"port"`
	Host    string `yaml:"host"`
	Timeout int64  `yaml:"timeout"`
}

func GetConfig(yamlName string) *Config {
	data, err := os.ReadFile(yamlName)
	if err != nil {
		log.Printf("error while trying to read %v: %v", yamlName, err)
	}
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Printf("error while unmarshalling %v: %v", yamlName, err)
	}
	return c
}
