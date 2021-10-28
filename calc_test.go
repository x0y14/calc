package calc

import (
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		{
			"1+2",
			3,
		},
		{
			"1*2+3*4*5",
			62,
		},
		{
			"(4+1)-5",
			0,
		},
		{
			"52-20-20",
			12,
		},
		{
			"3*232*(0-13)+323-3232/32",
			-8826,
		},
		{
			"(3+5)/2",
			4,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			calculator := NewCalculator(50)
			actual := calculator.Calculate(test.in)
			if actual != test.want {
				t.Fatalf("%v = %v, but result was %v", test.in, test.want, actual)
			}
		})
	}
}
