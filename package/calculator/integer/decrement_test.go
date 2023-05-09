package integer_test

import (
	"testing"

	"dinero.go/package/calculator/integer"
)

func TestDecrement(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		input  int
		expect int
	}

	tests := []test{
		{input: 2, expect: 1},
		{input: -2, expect: -3},
	}

	for _, tc := range tests {
		got := c.Decrement(tc.input)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
