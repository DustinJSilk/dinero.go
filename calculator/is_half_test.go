package calculator_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestIsHalf(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		value  int
		total  int
		expect bool
	}

	tests := []test{
		{value: 5, total: 10, expect: true},
		{value: -5, total: 10, expect: true},
		{value: 2, total: 10, expect: false},
		{value: 0, total: 10, expect: false},
		{value: 5, total: 0, expect: false},
	}

	for _, tc := range tests {
		got := calculator.IsHalf(tc.value, tc.total)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
