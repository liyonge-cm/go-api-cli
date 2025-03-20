package cmd

import (
	"fmt"

	"github.com/liyonge-cm/go-api-cli/cmd/option"
)

const (
	// VERSION is the current go-api-cli version.
	VERSION = "v1.0.7"
)

type CmdVersion struct{}

func (c *CmdVersion) CMD() string {
	return "version"
}

func (c *CmdVersion) Register(cmdMap map[string]Command) {
	cmdMap[c.CMD()] = c
}

func (c *CmdVersion) Help() string {
	return "get current go-api-cli version"
}

func (c *CmdVersion) SetArgs(args []string, option map[string]string) error {
	return nil
}

func (c *CmdVersion) Run() {
	fmt.Println("version:", VERSION)
}

func (c *CmdVersion) Sub() map[string]option.CommandOption {
	return nil
}
