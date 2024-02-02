package cmd

import (
	"flag"
	"fmt"
	"path"

	"github.com/liyonge-cm/go-api-cli/gen_frame"
)

type CmdInit struct {
	prjName  *string
	jsonCase *string
}

func (c *CmdInit) Get() {
	prjName := flag.String("init", "", "init project name")
	jsonCase := flag.String("j", "snake", "json case: camel or snake")
	c.prjName = prjName
	c.jsonCase = jsonCase
}

func (c *CmdInit) Run() {
	if c.prjName == nil || *c.prjName == "" {
		return
	}
	fmt.Println("init project", *c.prjName)
	outPath := path.Dir(*c.prjName)
	name := path.Base(*c.prjName)
	s := gen_frame.NewGenFrameService(&gen_frame.GenFrameConfig{
		OutPath:  outPath,
		PrjName:  name,
		JsonCase: *c.jsonCase,
	})
	err := s.GenFrame()
	if err != nil {
		fmt.Println("init project err,", err.Error())
		return
	}
	fmt.Println("init", *c.prjName, "success")
}
