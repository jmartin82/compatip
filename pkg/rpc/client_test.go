package rpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"

	"testing"
)

type grpcAdapter struct {
}

func (ga *grpcAdapter) Dial(target string, opts ...grpc.DialOption) (Connector, error) {
	return grpc.Dial(target, opts...)
}


type mockConnector struct {
	error bool
	result string
}

func (c mockConnector) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	if c.error {
		return  errors.New("Invoke error")
	}

	reply = &c.result
	return nil
}


type mockDialer struct {
	error bool
	connector Connector
}


func (d mockDialer) Dial(target string, opts ...grpc.DialOption) (Connector, error) {
	if d.error {
		return d.connector, errors.New("Dial error")
	}

	return d.connector, nil
}




func TestGRPCTransport_DoCall(t *testing.T) {
	type fields struct {
		client Dialer
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{

		{
			"Invalid url format",
			fields{
				client:&grpcAdapter{},
			},
			args{
				url:"aaa",
			},
			"",
			true,
		},
		{
			"Endpoint error",
			fields{
				client:mockDialer{
					error:true,
				},
			},
			args{
				url:"grpc://example.com",
			},
			"",
			true,
		},
		{
			"Invalid server response",
			fields{
				client:mockDialer{
					false,
					mockConnector{
						true,
						"",

					},
				},
			},
			args{
				url:"grpc://example.com:9000",
			},
			"",
			true,
		},
		{
			"Valid server response",
			fields{
				client:mockDialer{
					false,
					mockConnector{
						false,
						"",

					},
				},
			},
			args{
				url:"grpc://example.com",
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := NewGRPCTransport(tt.fields.client)
			got, err := re.DoCall(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCTransport.DoCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GRPCTransport.DoCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
