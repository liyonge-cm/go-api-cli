package main

import (
	"flag"
	"fmt"
	"os"

	"go-cli/config"
	"go-cli/gen"
	"go-cli/gen_frame"
)

func main() {
	genMod := flag.String("g", "api", "api or frame")
	configFile := flag.String("c", "./config/config.yml", "yml config")
	flag.Parse()
	if *genMod == "" {
		fmt.Println("invalid genMod")
		os.Exit(0)
	}
	if *configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	cfg := config.LoadConfig(*configFile)

	if *genMod == "api" {
		s := gen.NewGenServer()
		// 获取表字段
		s.ConnDB()
		s.GetTableFields()
		// 生成model
		s.GenModel()
		// 生成API
		s.GenApi()
	} else if *genMod == "frame" {
		s := gen_frame.NewGenFrameService(&gen_frame.GenFrameConfig{
			OutPath:  cfg.Frame.OutPath,
			PrjName:  cfg.Frame.PrjName,
			JsonCase: cfg.Frame.JsonCase,
		})
		err := s.GenFrame()
		if err != nil {
			fmt.Println("gen err,", err.Error())
		}
	}
}
