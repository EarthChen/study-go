package main

import (
	"testing"
)

func Test_Hello(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		inputLanguage string
		want      string
	}{
		// TODO: Add test cases.
		{
			"test1",
			"world",
			"Spanish",
			"Hola, world",
		},
		{
			"test2",
			"",
			"",
			"hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.inputName,tt.inputLanguage); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
