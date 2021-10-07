package concurrency

import (
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {
	type args struct {
		wc   WebsiteChecker
		urls []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				wc: mockWebsiteChecker,
				urls: []string{
					"aaa",
					"bbb",
					"ccc",
				},
			},
			want: map[string]bool{
				"aaa": false,
				"bbb": true,
				"ccc": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWebsites(tt.args.wc, tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWebsites() = %v, want %v", got, tt.want)
			}
		})
	}
}
