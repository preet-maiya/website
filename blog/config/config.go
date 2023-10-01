package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIGFILE = "config.yaml"

type Config struct {
	ResumeFile string `yaml:"resume_file"`
}

func GetConfig() *Config {
	c := &Config{}
	c.ResumeFile = "/resume.pdf"

	if _, err := os.Stat(CONFIGFILE); errors.Is(err, os.ErrNotExist) {
		return c
	}

	configFile, err := os.ReadFile(CONFIGFILE)
	if err != nil {
		log.Fatalf("Error reading yaml file %v with error %v", configFile, err)
	}

	if err = yaml.Unmarshal(configFile, c); err != nil {
		log.Fatalf("Error unmarshalling: %v", err)
	}

	return c
}
