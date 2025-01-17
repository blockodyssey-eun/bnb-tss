package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Kubernetes struct {
		Namespace       string `yaml:"namespace"`
		PodPrefix       string `yaml:"podPrefix"`
		PoImage         string `yaml:"podImage"`
		InitialPodCount int    `yaml:"initialPodCount"`
	} `yaml:"kubernetes"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
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
