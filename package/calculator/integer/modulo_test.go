package integer_test

import (
	"testing"

	"dinero.go/package/calculator/integer"
)

func TestModulo(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a         int
		b         int
		expect    int
		expectErr bool
	}

	tests := []test{
		{a: 5, b: 3, expect: 2},
		{a: -5, b: -4, expect: -1},
		{a: 4e5, b: 3e5, expect: 100000},
		{a: 1, b: 0, expectErr: true},
	}

	for _, tc := range tests {
		got, err := c.Modulo(tc.a, tc.b)
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
