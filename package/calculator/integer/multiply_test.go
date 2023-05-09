package integer_test

import (
	"testing"

	"dinero.go/package/calculator/integer"
)

func TestMultiply(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect int
	}

	tests := []test{
		{a: 10, b: 20, expect: 200},
		{a: -10, b: -20, expect: 200},
		{a: 1e5, b: 2e5, expect: 20000000000},
	}

	for _, tc := range tests {
		got := c.Multiply(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
