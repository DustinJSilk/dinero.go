package integer_test

import (
	"testing"

	"dinero.go/package/calculator/integer"
)

func TestIntegerDivide(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a         int
		b         int
		expect    int
		expectErr bool
	}

	tests := []test{
		// Divides positive numbers.
		{a: 8, b: 2, expect: 4},
		// Divides negative numbers.
		{a: -8, b: -2, expect: 4},
		// Divides numbers in scientific notation.
		{a: 3e5, b: 2e5, expect: 1},
		// Rounds positive numbers towards zero.
		{a: 3, b: 2, expect: 1},
		// Rounds negative numbers towards zero.
		{a: -3, b: 2, expect: -1},
		// Returns an error when dividing by zero.
		{a: 1, b: 0, expectErr: true},
	}

	for _, tc := range tests {
		got, err := c.IntegerDivide(tc.a, tc.b)
		if err != nil {
			if tc.expectErr {
				continue
			}
			t.Fatalf("error: %v, got: %v", tc.expect, got)
		}

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
