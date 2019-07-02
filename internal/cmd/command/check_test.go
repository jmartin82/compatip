package command

import (
	"flag"
	"github.com/jmartin82/compatip/pkg/rest"
	"testing"

	"github.com/codegangsta/cli"
)



func getCheckCommandContext(args ...string)  *cli.Context{
	app:=&cli.App{
		Name:"test",
	}



	flags:=flag.NewFlagSet("flags",1)
	flags.String("jsonpath","","jsonpath")
	flags.Parse(args)

	context:=cli.NewContext(app, flags,nil)
	context.Command.Action = "check"

	return context

}


func TestCmdCheck(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		client  *httpGetter
		args    args
		wantErr bool
	}{

		{
			"Invalid endpoint",
			&httpGetter{false,404,""},
			args{
				getCheckCommandContext("http://example.com/XXX"),
			},
			true,
		},
		{
			"Get invalid version",
			&httpGetter{false,200, "xxx"},
			args{
				getCheckCommandContext("http://example.com/version"),
			},
			true,
		},
		{
			"Get Version",
			&httpGetter{false,200, "v1.2.3"},
			args{
				getCheckCommandContext("http://example.com/version"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient = func() rest.HTTPGetter {
				return tt.client
			}
			if err := CmdCheck(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CmdCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
