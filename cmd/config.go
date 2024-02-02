package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/liyonge-cm/go-api-cli/config"
)

type CmdConfig struct {
	configFile *string
}

func (c *CmdConfig) Get() {
	configFile := flag.String("c", "./config/config.yml", "yml config")
	c.configFile = configFile
}

func (c *CmdConfig) Run() {
	if *c.configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	_ = config.LoadConfig(*c.configFile)
}
