package version

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Invalid version",
			args{
				"x",
			},
			"",
			true,
		},
		{
			"Valid version",
			args{
				"v1.2",
			},
			"1.2.0",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Check(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
