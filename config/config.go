package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "mars.yml"

var globalConfig = &Config{}

func Global() *Config {
	return globalConfig
}

func Init() {
	b, err := os.ReadFile(configFileName)
	if err != nil {
		panic(fmt.Sprintf("init mars failure, read file <%s> error: %v", configFileName, err))
	}

	err = yaml.Unmarshal(b, globalConfig)
	if err != nil {
		panic(fmt.Sprintf("init mars failure, unmarshal file <%s> error: %v", configFileName, err))
	}
}

type Host struct {
	Domain   string `yaml:"domain"`
	Port     int    `yaml:"port"`
	UseSSL   int    `yaml:"use_ssl"`
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

type Cookie struct {
	ExpireHours int `yaml:"expire_hours"`
	HttpOnly    int `yaml:"http_only"`
}

type Config struct {
	Host   Host   `yaml:"host"`
	Cookie Cookie `yaml:"cookie"`
}
