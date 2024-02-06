package config

import (
	"log"
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

func LoadConfig(filename string) {
	if filename == "" {
		filename = "config.yml"
	}
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// 解析YAML文件
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
