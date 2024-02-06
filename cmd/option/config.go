package option

import (
	"fmt"
	"os"

	"github.com/liyonge-cm/go-api-cli/config"
)

type OptConfig struct {
	configFile string
}

func (c *OptConfig) Run() {
	if c.configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	_, err := config.LoadConfig(c.configFile)
	if err != nil {
		fmt.Println("load config file error:", err)
		os.Exit(0)
	}
}

func (c *OptConfig) OPTION() string {
	return "c"
}

func (c *OptConfig) Help() string {
	return "use config file to generate api, default: ./config/config.yml"
}

func (c *OptConfig) SetOptions(option string) error {
	// 默认
	if option == "" {
		option = "./config/config.yml"
	}

	c.configFile = option
	return nil
}

func (c *OptConfig) Get() string {
	return c.configFile
}
