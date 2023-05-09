package integer_test

import (
	"testing"

	"dinero.go/calculator/integer"
)

func TestPowInt(t *testing.T) {
	type test struct {
		a      int
		b      int
		expect int
	}

	tests := []test{
		{a: 5, b: 1, expect: 5},
		{a: 5, b: 0, expect: 1},
		{a: 2, b: 3, expect: 8},
		{a: -2, b: 3, expect: -8},
		{a: 123, b: 7, expect: 425927596977747},
		{a: 1e5, b: 3, expect: 1000000000000000},
	}

	for _, tc := range tests {
		got := integer.PowInt(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
