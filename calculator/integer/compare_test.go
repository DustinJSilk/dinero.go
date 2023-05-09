package integer_test

import (
	"testing"

	"dinero.go/calculator/integer"
	"dinero.go/types"
)

func TestCompare(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect types.CompareResult
	}

	tests := []test{
		{a: 1, b: 2, expect: types.LT},
		{a: -2, b: -1, expect: types.LT},
		{a: 1e5, b: 2e5, expect: types.LT},
		{a: 2, b: 1, expect: types.GT},
		{a: -1, b: -2, expect: types.GT},
		{a: 2e5, b: 1e5, expect: types.GT},
		{a: 2, b: 2, expect: types.EQ},
		{a: -2, b: -2, expect: types.EQ},
		{a: 2e5, b: 2e5, expect: types.EQ},
	}

	for _, tc := range tests {
		got := c.Compare(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
