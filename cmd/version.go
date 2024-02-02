package cmd

import (
	"flag"
	"fmt"
)

const (
	// VERSION is the current go-api-cli version.
	VERSION = "v1.0.0"
)

type CmdVersion struct {
	version *bool
}

func (c *CmdVersion) Get() {
	version := flag.Bool("v", false, "version")
	c.version = version
}

func (c *CmdVersion) Run() {
	if *c.version {
		fmt.Println("version:", VERSION)
	}
}
