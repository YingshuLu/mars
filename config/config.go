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

func init() {
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
	MaxAge   int `yaml:"max_age"`
	HttpOnly int `yaml:"http_only"`
	Secure   int `yaml:"secure"`
}

type Signer struct {
	Name    string `yaml:"name"`
	KeyPath string `yaml:"key_path"`
}

type Security struct {
	Signer    Signer `yaml:"signer"`
	Validator Signer `yaml:"validator"`
}

type Badger struct {
	ProfileDB string `yaml:"profile_path"`
}

type Storage struct {
	Badger Badger `yaml:"badger"`
}

type Config struct {
	Host     Host     `yaml:"host"`
	Cookie   Cookie   `yaml:"cookie"`
	Security Security `yaml:"security"`
	Storage  Storage  `yaml:"storage"`
}
