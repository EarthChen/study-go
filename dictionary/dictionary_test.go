package dictionary

import "testing"

func TestDict_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		d       Dict
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			d: Dict{
				"string": "string",
			},
			args: args{
				word: "1",
			},
			wantErr: true,
		},
		{
			name: "test2",
			d: Dict{
				"string": "string",
			},
			args: args{
				word: "string",
			},
			want:    "string",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Search(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
