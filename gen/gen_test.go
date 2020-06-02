package gen

import "testing"

func Test_ok(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "5--1",
			args: args{
				username: "5--1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUsername(tt.args.username); got != tt.want {
				t.Errorf("validUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
