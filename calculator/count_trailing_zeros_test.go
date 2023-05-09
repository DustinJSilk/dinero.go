package calculator_test

import (
	"testing"

	"dinero.go/calculator"
	"dinero.go/calculator/integer"
)

func TestCountTrailingZeros(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		value     int
		base      int
		expect    int
		expectErr bool
	}

	tests := []test{
		{value: 1000, base: 10, expect: 3},
		{value: -1000, base: 10, expect: 3},
		{value: 1e3, base: 10, expect: 3},
		{value: -1e3, base: 10, expect: 3},
		{value: 123, base: 10, expect: 0},
		{value: -123, base: 10, expect: 0},
		{value: 0, base: 10, expect: 0},
		{value: 0, base: 2, expect: 0},
		{value: 123, base: 0, expectErr: true},
	}

	for _, tc := range tests {
		got, err := c.CountTrailingZeros(tc.value, tc.base)
		if err != nil {
			if tc.expectErr {
				continue
			}
			t.Fatalf("error: %v: %v, got: %v", err, tc.expect, got)
		}

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
