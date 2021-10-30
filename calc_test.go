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
		{
			"+2*40",
			80,
		},
		{
			"-80*3-20",
			-260,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			calculator := NewCalculator(50)
			result := calculator.Calculate(test.in)
			if result != test.want {
				t.Fatalf("%v = %v, but result was %v", test.in, test.want, result)
			}
		})
	}
}

func TestCalculator_Is(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		{
			"0 < 5",
			true,
		},
		{
			"0 > 5",
			false,
		},
		{
			"3*232*(0-13)+323-3232/32 < 0",
			true,
		},
		{
			"-1 * (3*232*(0-13)+323-3232/32) == -(3*232*(0-13)+323-3232/32)",
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			calculator := NewCalculator(50)
			result := calculator.Is(test.in)
			if result != test.want {
				t.Fatalf("%v = %v, but result was %v", test.in, test.want, result)
			}
		})
	}
}
