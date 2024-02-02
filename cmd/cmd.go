package cmd

import "flag"

type Command interface {
	Get()
	Run()
}

type Cmd struct {
	cmds []Command
}

func NewCommand() *Cmd {
	return &Cmd{
		cmds: []Command{
			&CmdVersion{},
			&CmdInit{},
			&CmdGen{
				cfg: &CmdConfig{},
			},
		},
	}
}

func (c *Cmd) Run() {
	for _, cmd := range c.cmds {
		cmd.Get()
	}
	flag.Parse()
	for _, cmd := range c.cmds {
		cmd.Run()
	}
}
