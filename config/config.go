package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg Config

type Config struct {
	Frame Frame
	Api   Api
	MySQL MySQL
}

type Frame struct {
	OutPath  string `yaml:"out_path"`
	PrjName  string `yaml:"prj_name"`
	JsonCase string `yaml:"json_case"`
}

type Api struct {
	Tables []string `yaml:"tables"`
}

type MySQL struct {
	Endpoint     string `yaml:"endpoint"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func LoadConfig(filename string) (Config, error) {
	if filename == "" {
		filename = "config.yml"
	}
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return Cfg, err
	}

	// 解析YAML文件
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		return Cfg, err
	}
	return Cfg, nil
}
