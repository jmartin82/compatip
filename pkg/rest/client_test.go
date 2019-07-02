package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type httpGetter struct {
	error bool
	status int
	body []byte
}

func (h httpGetter) Get(string) (resp *http.Response, err error) {
	if h.error {
		return &http.Response{}, errors.New("some error")
	}

	return &http.Response{
		StatusCode:h.status,
		Body:ioutil.NopCloser(bytes.NewReader(h.body)) ,
	}, nil
}

func TestRESTConnector_DoCall(t *testing.T) {
	type fields struct {
		client HTTPGetter
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
			"Invalid url",
			fields{
				client:http.DefaultClient,
			},
			args{
				url:"test",
			},
			"",
			true,
		},
		{
			"Call error",
			fields{
				client:httpGetter{
					error:true,
				},
			},
			args{
				url:"http://example.com",
			},
			"",
			true,
		},
		{
			"Invalid server response",
			fields{
				client:httpGetter{
					error:false,
					body:nil,
					status:400,

				},
			},
			args{
				url:"http://example.com",
			},
			"",
			true,
		},
		{
			"Valid call call",
			fields{
				client:httpGetter{
				error:false,
				body:[]byte("1.2.3"),
				status:200,

			},
			},
			args{
				url:"http://example.com",
			},
			"1.2.3",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := NewRESTTransport(tt.fields.client)
			got, err := re.DoCall(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("RESTTransport.DoCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RESTTransport.DoCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
