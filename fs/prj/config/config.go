package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg Config

type Config struct {
	Service Service
	MySQL   MySQL
}

type Service struct {
	Port int `yaml:"port"`
}

type MySQL struct {
	Endpoint     string `yaml:"endpoint"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func LoadConfig(filename string) (cfg Config, err error) {
	if filename == "" {
		filename = "config.yml"
	}
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return cfg, err
	}

	// 解析YAML文件
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return cfg, err
	}

	Cfg = cfg
	return cfg, err
}
