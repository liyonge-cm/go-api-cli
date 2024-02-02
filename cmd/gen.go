package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/liyonge-cm/go-api-cli/gen"
	"github.com/liyonge-cm/go-api-cli/gen_frame"
)

type CmdGen struct {
	mod   *string
	table *string
	cfg   *CmdConfig
}

func (c *CmdGen) Get() {
	c.cfg.Get()
	mod := flag.String("g", "", "generating module: api")
	table := flag.String("t", "", "generating api with table: table name")

	c.mod = mod
	c.table = table
}

func (c *CmdGen) Run() {
	if c.mod == nil || *c.mod == "" {
		return
	}
	fmt.Println("gen", *c.mod, "......")
	if *c.table != "" {
		fmt.Println("table:", *c.table)
	}
	c.cfg.Run()
	switch *c.mod {
	case "api":
		var tables []string
		if *c.table != "" {
			tables = strings.Split(*c.table, ",")
		}
		s := gen.NewGenServer(tables)
		// 获取表字段
		if err := s.ConnDB(); err != nil {
			fmt.Println("gen", *c.mod, "failed", "connect db err")
			return
		}
		if err := s.GetTableFields(); err != nil {
			fmt.Println("gen", *c.mod, "failed", "get table fields err")
			return
		}
		// 生成model
		if err := s.GenModel(); err != nil {
			fmt.Println("gen", *c.mod, "failed", "generate go model err")
			return
		}
		// 生成API
		if err := s.GenApi(); err != nil {
			fmt.Println("gen", *c.mod, "failed", "generate api err", err.Error())
			return
		}
		fmt.Println("gen", *c.mod, "api success")

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
