package rpc

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/grpc"
)

type Dialer interface {
	Dial(target string, opts ...grpc.DialOption) (Connector, error)
}

type Connector interface {
	Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error
}

type GRPCTransport struct {
	client Dialer
}

func NewGRPCTransport(client Dialer) *GRPCTransport {
	return &GRPCTransport{client: client}
}

func (re *GRPCTransport) DoCall(u string) (string, error) {

	url, err := url.ParseRequestURI(u)
	if err != nil {
		return "", err
	}

	server:= fmt.Sprintf("%s:%s",url.Hostname(), func(port string) string{
		if port==""{
			return "80"
		}
		return port
	} (url.Port()))

	conn, err := re.client.Dial(server, grpc.WithInsecure())
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	opts := []grpc.CallOption{}

	in := &Empty{}
	out := &VersionMessage{}

	err = conn.Invoke(ctx, url.RequestURI() , in, out, opts...)
	if err != nil {
		return "", err
	}

	return out.GetVersion(), nil
}
