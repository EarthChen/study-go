package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1111",
			args: args{
				"1",
			},
			want: "11111",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.char); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
