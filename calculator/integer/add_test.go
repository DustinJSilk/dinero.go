package integer_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestAdd(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect int
	}

	tests := []test{
		{a: 2, b: 3, expect: 5},
		{a: -1, b: -2, expect: -3},
		{a: 1e5, b: 2e5, expect: 300000},
	}

	for _, tc := range tests {
		got := c.Add(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
