package command

import (
	"errors"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jmartin82/compatip/pkg/version"
	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
)

func CmdAssert(c *cli.Context) error {

	//arguments
	if c.NArg() < 2 {
		return commandError(errors.New("Invalid number of arguments."), 2)
	}

	//banner
	commandInfo(c)

	//extraction
	url := c.Args().First()
	constraint := strings.Join(c.Args().Tail(), ",")
	v, err := extractVersion(url, c)
	if err != nil {
		return commandError(err, 3)
	}
	emoji.Printf("%s %s\n", ":goat: Version:", Green(v).Bold())

	//check
	res, err := version.Assert(v, constraint)
	if err != nil {
		return commandError(err, 4)
	}
	if res {
		emoji.Printf("%s\n", ":rocket: Compatible!")
	} else {
		return commandError(errors.New("Incompatible"), 1)
	}

	return nil
}
