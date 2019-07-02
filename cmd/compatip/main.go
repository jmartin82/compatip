package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jmartin82/compatip/internal/cmd/command"
)

const Name string = "compatip"
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var CommonFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "jsonpath",
		Value: "",
		Usage: "Dot notation path to get the version",
	},
}

var Commands = []cli.Command{
	{
		Name:   "check",
		Usage:  "Checks the current version of the service",
		Action: command.CmdCheck,
		Flags:  CommonFlags,
	},
	{
		Name:   "assert",
		Usage:  "Asserts the current version of the service with the specified constraints",
		Action: command.CmdAssert,
		Flags:  CommonFlags,
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = version
	app.Author = "jmartin82"
	app.Email = "jordimartin@gmail.com"

	app.Usage = "it's a simple tool to ensure compatibility between microservices."

	//app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)

}
