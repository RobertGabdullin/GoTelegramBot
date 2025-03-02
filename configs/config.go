package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type BotConfig struct {
	Token         string `yaml:"token"`
	UpdateTimeout int    `yaml:"update_timeout"`
	UpdateOffset  int    `yaml:"update_offset"`
}

type ScrapperConfig struct {
	BaseUrl        string `yaml:"base_url"`
	UpdateInterval int    `yaml:"scrapper_update_interval"`
	DatabaseUrl    string `yaml:"database_url"`
}

type Config struct {
	Bot      BotConfig      `yaml:"bot"`
	Scrapper ScrapperConfig `yaml:"scrapper"`
}

func LoadConfig() (*Config, error) {
	configFile, err := os.Open("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
