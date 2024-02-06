package cmd

import (
	"errors"
	"fmt"
	"path"

	"github.com/liyonge-cm/go-api-cli/cmd/option"
	"github.com/liyonge-cm/go-api-cli/gen_frame"
)

type CmdInit struct {
	prjName  string
	jsonCase option.CommandOption
}

func (c *CmdInit) CMD() string {
	return "init"
}

func (c *CmdInit) Register(cmdMap map[string]Command) {
	cmdMap[c.CMD()] = c

	c.jsonCase = &option.OptJsonCase{}
}

func (c *CmdInit) Help() string {
	return "init a project, argument: project name"
}

func (c *CmdInit) SetArgs(args []string, option map[string]string) error {
	if len(args) == 0 {
		return errors.New("init project name is required")
	}
	c.prjName = args[0]
	if option == nil {
		return nil
	}

	if err := c.jsonCase.SetOptions(option[c.jsonCase.OPTION()]); err != nil {
		return err
	}

	return nil
}

func (c *CmdInit) Run() {
	fmt.Println("init project", c.prjName)
	if c.prjName == "" {
		fmt.Println("init project name is required")
	}

	outPath := path.Dir(c.prjName)
	name := path.Base(c.prjName)

	s := gen_frame.NewGenFrameService(&gen_frame.GenFrameConfig{
		OutPath:  outPath,
		PrjName:  name,
		JsonCase: c.jsonCase.Get(),
	})
	err := s.GenFrame()
	if err != nil {
		fmt.Println("init project err,", err.Error())
		return
	}
	fmt.Println("init", c.prjName, "success")
}

func (c *CmdInit) Sub() map[string]option.CommandOption {
	return map[string]option.CommandOption{
		c.jsonCase.OPTION(): c.jsonCase,
	}
}
