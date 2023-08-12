package config

import (
	"gopkg.in/yaml.v3"
	"log"
)

type Database struct {
	Driver       string `yaml:"driver"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
}

type Config struct {
	Database Database `yaml:"database"`
}

func LoadConfig(configFileData []byte) (*Config, error) {
	config := new(Config)

	if err := yaml.Unmarshal(configFileData, &config); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
		return config, err
	}

	return config, nil
}
