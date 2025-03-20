package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/liyonge-cm/go-api-cli/cmd/option"
	"github.com/liyonge-cm/go-api-cli/gen"
	"github.com/liyonge-cm/go-api-cli/gen_frame"
)

type CmdGen struct {
	mod   string
	table option.OptTable
	cfg   option.OptConfig
}

func (c *CmdGen) CMD() string {
	return "gen"
}

func (c *CmdGen) Register(cmdMap map[string]Command) {
	cmdMap[c.CMD()] = c
}

func (c *CmdGen) Help() string {
	return "generating module: api or model"
}

func (c *CmdGen) SetArgs(args []string, option map[string]string) error {
	if len(args) == 0 {
		return errors.New("generating module is required, api or model")
	}
	c.mod = args[0]
	if c.mod != "api" && c.mod != "model" {
		return errors.New("generating module is error, api or model")
	}
	if option == nil {
		return nil
	}

	if err := c.table.SetOptions(option[c.table.OPTION()]); err != nil {
		return err
	}
	if err := c.cfg.SetOptions(option[c.cfg.OPTION()]); err != nil {
		return err
	}

	c.mod = args[0]
	return nil
}

func (c *CmdGen) Run() {
	if c.mod == "" {
		return
	}
	fmt.Println("gen", c.mod, "......")

	table := c.table.Get()
	if table != "" {
		fmt.Println("table:", table)
	}
	c.cfg.Run()
	switch c.mod {
	case "api":
		var tables []string
		if table != "" {
			tables = strings.Split(table, ",")
		}
		s := gen.NewGenServer(tables)
		// 获取表字段
		if err := s.ConnDB(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "connect db err,", err.Error())
			return
		}
		if err := s.GetTableFields(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "get table fields err,", err.Error())
			return
		}
		// 生成model
		if err := s.GenModel(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "generate go model err,", err.Error())
			return
		}
		// 生成API
		if err := s.GenApi(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "generate api err,", err.Error())
			return
		}
		fmt.Println("gen", c.mod, "success")

	case "model":
		var tables []string
		if table != "" {
			tables = strings.Split(table, ",")
		}
		s := gen.NewGenServer(tables)
		// 获取表字段
		if err := s.ConnDB(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "connect db err,", err.Error())
			return
		}
		if err := s.GetTableFields(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "get table fields err,", err.Error())
			return
		}
		// 生成model
		if err := s.GenModel(); err != nil {
			fmt.Println("gen", c.mod, "failed,", "generate go model err,", err.Error())
			return
		}

	case "frame":
		s := gen_frame.NewGenFrameService(nil)
		err := s.GenFrame()
		if err != nil {
			fmt.Println("gen frame err,", err.Error())
		}
	default:
		fmt.Println("invalid gen mod")
	}
}

func (c *CmdGen) Sub() map[string]option.CommandOption {
	return map[string]option.CommandOption{
		c.table.OPTION(): &c.table,
		c.cfg.OPTION():   &c.cfg,
	}
}
