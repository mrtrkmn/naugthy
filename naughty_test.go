package naughty

import (
	"testing"
)

func TestIsNaugthy(t *testing.T) {

	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "Not naughty word 1", word: "1234567985234", want: false},
		{name: "Not naughty word 2", word: "testcasenaughty", want: false},
		{name: "Naughty word 1 ", word: "ZSB0aXRsZT0+", want: true},
		{name: "Naughty word 1 ", word: "2q/ahtm+2pg=", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNaugthy(tt.word); got != tt.want {
				t.Errorf("IsNaugthy() = %v, want %v", got, tt.want)
			}
		})
	}
}
