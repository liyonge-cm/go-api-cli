package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/liyonge-cm/go-api-cli-prj/app"
	"github.com/liyonge-cm/go-api-cli-prj/config"
)

func main() {
	configFile := flag.String("c", "./config/config.yml", "yml config")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println("load config fail", err.Error())
		os.Exit(0)
	}
	app := app.NewApp(&cfg)
	app.Run()
}
