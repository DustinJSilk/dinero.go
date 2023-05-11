package calculator_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestAbsolute(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		input  int
		expect int
	}

	tests := []test{
		{input: 5, expect: 5},
		{input: -5, expect: 5},
		{input: 0, expect: 0},
		{input: -0, expect: 0},
	}

	for _, tc := range tests {
		got := c.Absolute(tc.input)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
