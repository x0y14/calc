package calc

import (
	"fmt"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	var tests = []struct {
		title string
		in    string
	}{
		{
			"simple",
			"1+12",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tk := NewTokenizer(test.in)
			tokens, err := tk.Tokenize()
			if err != nil {
				t.Fatal(err)
			}

			for i, tok := range tokens {
				fmt.Printf("%v) %v\n", i, tok.String())
			}

		})
	}
}
