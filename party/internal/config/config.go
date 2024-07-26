package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GRPC struct {
		Port int `yaml:"port"`
	} `yaml:"grpc"`
	Gateway struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"gateway"`
}

var cfg Config

func Load() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return fmt.Errorf("error opening config file: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return fmt.Errorf("error decoding config file: %v", err)
	}

	return nil
}

func Get() *Config {
	return &cfg
}
