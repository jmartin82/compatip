package version

import (
	"errors"
	"testing"
)

type MockTransport struct {
	error bool
	content string
}

func (mt MockTransport) DoCall(string) (string, error)  {
	if mt.error {
		return "", errors.New("Somthing wrong")
	}
	return mt.content, nil
}



func TestExtractor_Extract(t *testing.T) {
	type fields struct {
		url       string
		transport Transport
		post      PostProcessor
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"Invalid URL",
			fields{
				url: "invalid url",
			},
			"",
			true,
		},
		{
			"Error on transport",
			fields{
				url: "http://example.com",
				transport: MockTransport{
					error:true,
				},
			},
			"",
			true,
		},
		{
			"Correct extraction without post process",
			fields{
				url: "http://example.com",
				post:nil,
				transport: MockTransport{
					error:false,
					content: "{ \"vvv\":\"1234\" }",
				},
			},
			"{ \"vvv\":\"1234\" }",
			false,
		},
		{
			"Correct extraction with process",
			fields{
				url: "http://example.com",
				post:NewJsonPostProcessor("vvv"),
				transport: MockTransport{
					error:false,
					content: "{ \"vvv\":\"1234\" }",
				},
			},
			"1234",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewExtractor()
			e.WithTransport(tt.fields.transport)
			e.WithPostProcessor(tt.fields.post)
			e.FromURL(tt.fields.url)
			got, err := e.Extract()
			if (err != nil) != tt.wantErr {
				t.Errorf("Extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Extractor.Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
