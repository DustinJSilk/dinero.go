package calculator_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestMinimum(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		{input: []int{5, 3, 2}, expect: 2},
		{input: []int{-5, -4, -2}, expect: -5},
		{input: []int{4e5, 3e5, 2e5}, expect: 2e5},
	}

	for _, tc := range tests {
		got := calculator.Minimum(tc.input...)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
