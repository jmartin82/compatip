package command

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/jmartin82/compatip/pkg/rest"
	"github.com/jmartin82/compatip/pkg/rpc"
	"github.com/jmartin82/compatip/pkg/version"
	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)


var httpClient = func() rest.HTTPGetter {
	return http.DefaultClient
}

type grpcAdapter struct {
}

func (ga *grpcAdapter) Dial(target string, opts ...grpc.DialOption) (rpc.Connector, error) {
	return grpc.Dial(target, opts...)
}

var rpcClient = func() *grpcAdapter {
	return &grpcAdapter{}
}


func commandInfo(c *cli.Context)  {

	if (c.NArg()>0) {
		url:= c.Args().First()
		emoji.Printf("%s %s\n", ":point_down: Checking:", Cyan(url))
		emoji.Printf("%s %s\n", ":wave: Transport:", Yellow(getSchema(url)))
	}


	if (c.NArg() >1 && c.Command.Name == "assert" ) {
		constrain := c.Args().Tail()
		emoji.Printf("%s %v\n", ":speak_no_evil: Constraint:", Green(constrain))
	}


}

func commandError(value error, code int) *cli.ExitError {
	msg := emoji.Sprintf(":shit: Error: %s\n", Red(value))
	return cli.NewExitError(msg, code)
}

func getSchema(url string) string {
	i:=strings.Index(url,"://")
	if i < 0{
		return "REST"
	}

	schema:= strings.ToLower(url[0:i])
	switch schema {
	case "http":
		fallthrough
	case "https":
		return "REST"
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "gproto+http":
		return "GRPC"
	default:
		return "UNKNOWN"
	}
}


func getTransport(schema string) (version.Transport, error) {
	var transport version.Transport
	switch strings.ToUpper(schema) {
	case "GRPC":
		transport = rpc.NewGRPCTransport(rpcClient())
	case "REST":
		transport = rest.NewRESTTransport(httpClient())
	default:
		msg := fmt.Sprintf("Invalid transport %s", schema)
		return nil, errors.New(msg)
	}
	return transport, nil
}




func extractVersion(url string, c *cli.Context) (string, error) {

	extractor := version.NewExtractor()
	extractor.FromURL(url)

	transport,err := getTransport(getSchema(url))
	if err!=nil {
		return  "",err
	}
	extractor.WithTransport(transport)

	if c.String("jsonpath") != "" {
		jpp := version.NewJsonPostProcessor(c.String("jsonpath"))
		extractor.WithPostProcessor(jpp)
	}
	v, err := extractor.Extract()
	return v, err
}

