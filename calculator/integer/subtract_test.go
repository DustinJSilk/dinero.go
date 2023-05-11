package integer_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestSubtract(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect int
	}

	tests := []test{
		{a: 1, b: 2, expect: -1},
		{a: -1, b: -2, expect: 1},
		{a: 1e5, b: 2e5, expect: -100000},
	}

	for _, tc := range tests {
		got := c.Subtract(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
