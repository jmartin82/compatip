package version

import "testing"

func TestAssert(t *testing.T) {
	type args struct {
		v          string
		constraint string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Empty version",
			args{
				"",
				"1.2.3",
			},
			false,
			true,
		},
		{
			"Empty constraint",
			args{
				"1.2.3",
				"",
			},
			false,
			true,
		},
		{
			"Invalid version",
			args{
				"x",
				"1.2.3",
			},
			false,
			true,
		},
		{
			"Invalid constraint",
			args{
				"1.2.3",
				"x",
			},
			false,
			true,
		},
		{
			"Invalid version",
			args{
				"1.2.3",
				">1.4.0",
			},
			false,
			false,
		},
		{
			"Valid version",
			args{
				"1.2.3",
				">1.2",
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Assert(tt.args.v, tt.args.constraint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Assert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Assert() = %v, want %v", got, tt.want)
			}
		})
	}
}
