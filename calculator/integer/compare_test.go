package integer_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestCompare(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect calculator.CompareResult
	}

	tests := []test{
		{a: 1, b: 2, expect: calculator.LT},
		{a: -2, b: -1, expect: calculator.LT},
		{a: 1e5, b: 2e5, expect: calculator.LT},
		{a: 2, b: 1, expect: calculator.GT},
		{a: -1, b: -2, expect: calculator.GT},
		{a: 2e5, b: 1e5, expect: calculator.GT},
		{a: 2, b: 2, expect: calculator.EQ},
		{a: -2, b: -2, expect: calculator.EQ},
		{a: 2e5, b: 2e5, expect: calculator.EQ},
	}

	for _, tc := range tests {
		got := c.Compare(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
