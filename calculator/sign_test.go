package calculator_test

import (
	"testing"

	"dinero.go/calculator"
	"dinero.go/calculator/integer"
)

func TestSign(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		input  int
		expect int
	}

	tests := []test{
		{input: 0, expect: 0},
		{input: -0, expect: 0},
		{input: 5, expect: 1},
		{input: -5, expect: -1},
	}

	for _, tc := range tests {
		got := c.Sign(tc.input)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
