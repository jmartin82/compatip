package version

import (
	"reflect"
	"testing"
)

func TestNewJsonPostProcessor(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *JsonPostProcessor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJsonPostProcessor(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJsonPostProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonPostProcessor_Process(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jp := &JsonPostProcessor{
				path: tt.fields.path,
			}
			if got := jp.Process(tt.args.in); got != tt.want {
				t.Errorf("JsonPostProcessor.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
