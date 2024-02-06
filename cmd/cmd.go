package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/liyonge-cm/go-api-cli/cmd/option"
)

type Command interface {
	CMD() string
	Register(map[string]Command)
	Help() string
	SetArgs([]string, map[string]string) error
	Run()
	Sub() map[string]option.CommandOption
}

// type CommandOption interface {
// 	OPTION() string
// 	Help() string
// 	SetOptions(string) error
// 	Get() string
// }

type Cmd struct {
	cmds   []Command
	cmdMap map[string]Command
}

func NewCommand() *Cmd {
	c := &Cmd{
		cmds: []Command{
			&CmdInit{},
			&CmdVersion{},
			&CmdGen{},
		},
		cmdMap: map[string]Command{},
	}
	for _, cmd := range c.cmds {
		cmd.Register(c.cmdMap)
	}
	return c
}

func (c *Cmd) Help() {
	fmt.Println(`
USAGE
	go-api-cli COMMAND [ARGUMENT] [OPTION]

COMMAND`)
	for _, cmd := range c.cmds {
		h := fmt.Sprintf("	%v	%v", cmd.CMD(), cmd.Help())
		fmt.Println(h)
		if cmd.Sub() != nil {
			for k, v := range cmd.Sub() {
				h := fmt.Sprintf("		%v	%v", k, v.Help())
				fmt.Println(h)
			}
		}
	}
}

func (c *Cmd) Run() {
	parsedArgs, parsedOptions := c.parse()
	if len(parsedArgs) <= 1 {
		c.Help()
		return
	}
	cmd := parsedArgs[1]
	if command, ok := c.cmdMap[cmd]; ok {
		if err := command.SetArgs(parsedArgs[2:], parsedOptions); err != nil {
			fmt.Println(err)
			return
		}
		command.Run()
	} else {
		c.Help()
	}
}

// ParseUsingDefaultAlgorithm parses arguments using default algorithm.
func (c *Cmd) parse() (parsedArgs []string, parsedOptions map[string]string) {
	args := os.Args
	argumentRegex := regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)

	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); {
		array := argumentRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" {
				parsedOptions[array[1]] = array[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					// Eg: gf gen -d -n 1
					parsedOptions[array[1]] = array[3]
				} else {
					// Eg: gf gen -n 2
					parsedOptions[array[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				// Eg: gf gen -h
				parsedOptions[array[1]] = array[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}
