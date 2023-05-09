package calculator_test

import (
	"testing"

	"dinero.go/calculator"
	"dinero.go/calculator/integer"
)

func TestIsEven(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		input  int
		expect bool
	}

	tests := []test{
		{input: 202, expect: true},
		{input: -202, expect: true},
		{input: 101, expect: false},
		{input: -101, expect: false},
		{input: 0, expect: true},
	}

	for _, tc := range tests {
		got := c.IsEven(tc.input)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
