package command

import (
	"errors"
	"flag"
	"github.com/jmartin82/compatip/pkg/rest"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/codegangsta/cli"
)



type httpGetter struct {
	error bool
	status int
	body string
}

func (h httpGetter) Get(string) (resp *http.Response, err error) {
	if h.error {
		return &http.Response{}, errors.New("some error")
	}

	return &http.Response{
		StatusCode:h.status,
		Body:ioutil.NopCloser(strings.NewReader(h.body)) ,
	}, nil
}




func getAssertCommandContext(args ...string)  *cli.Context{
	app:=&cli.App{
		Name:"test",
	}



	flags:=flag.NewFlagSet("flags",1)
	flags.String("jsonpath","","jsonpath")
	flags.Parse(args)

	context:=cli.NewContext(app, flags,nil)
	context.Command.Action = "assert"

	return context

}

func TestCmdAssert(t *testing.T) {


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
			"Invalid Arguments",
			&httpGetter{},
			args{
				getAssertCommandContext("one"),
			},
			true,
		},
		{
			"Invalid transport",
			&httpGetter{true,400,""},
			args{
				getAssertCommandContext("xxl://example.com","v1.2.3"),
			},
			true,
		},
		{
			"Invalid endpoint",
			&httpGetter{false,404,""},
			args{
				getAssertCommandContext("http://example.com/XXX","v1.2.3"),
			},
			true,
		},
		{
			"Invalid constrain",
			&httpGetter{false,200, "v1.2.3"},
			args{
				getAssertCommandContext("http://example.com/version","qqq"),
			},
			true,
		},
		{
			"Assert version",
			&httpGetter{false,200, "v1.2.3"},
			args{
				getAssertCommandContext("http://example.com/version","v1.2.3"),
			},
			false,
		},
		{
			"Assert Range",
			&httpGetter{false,200, "v1.4"},
			args{
				getAssertCommandContext("http://example.com/version",">v1.2.3"),
			},
			false,
		},
		{
			"Assert Range II",
			&httpGetter{false,200, "v1.4"},
			args{
				getAssertCommandContext("http://example.com/version",">v1.2.3"),
			},
			false,
		},
		{
			"Assert Range III (invalid)",
			&httpGetter{false,200, "v1.4"},
			args{
				getAssertCommandContext("http://example.com/version",">v1.2","<v1.3"),
			},
			true,
		},
		{
			"Json version extraction",
			&httpGetter{false,200, "{\"app\":{\"version\":\"v1.2.5\"}}"},
			args{
				getAssertCommandContext("--jsonpath=app.version","http://example.com/version",">v1.2","<v1.3"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient = func() rest.HTTPGetter {
				return tt.client
			}

			if err := CmdAssert(tt.args.c); (err != nil) != tt.wantErr {
				t.Log(err)
				t.Errorf("CmdAssert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
