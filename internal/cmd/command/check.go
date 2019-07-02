package command

import (
	"errors"
	"github.com/codegangsta/cli"
	"github.com/jmartin82/compatip/pkg/version"
	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
)

func CmdCheck(c *cli.Context) error {

	//arguments
	if c.NArg() == 0 {
		return commandError(errors.New("Invalid number of arguments."),2)
	}

	//banner
	commandInfo(c)

	//extraction
	url := c.Args().First()
	v, err := extractVersion(url, c)
	if err != nil {
		return commandError(err,3)
	}

	//check
	cv,err := version.Check(v)
	if err != nil {
		return commandError(err,4)
	}

	emoji.Printf("%s %s\n", ":goat: Version:", Green(cv).Bold())
	return nil
}

