package calc

import (
	"fmt"
	"testing"
)

func TestParser_Expr(t *testing.T) {
	var tests = []struct {
		title string
		in    string
	}{
		{
			"simple",
			"1+12",
		},
		{
			"1*2+(3+4)",
			"1*2+(3+4)",
		},
		{
			"1 == 1 && 2 == 2",
			"1 == 1 && 2 == 2;",
		},
		{
			"var",
			"ident = 1 + 2;",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			tk := NewTokenizer(test.in)
			tokens, err := tk.Tokenize()
			if err != nil {
				t.Fatal(err)
			}

			ps := NewParser(tokens)
			node, err := ps.Stmt()
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("%v\n", node)
		})
	}
}
