package integer_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestIncrement(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		input  int
		expect int
	}

	tests := []test{
		{input: 2, expect: 3},
		{input: -2, expect: -1},
	}

	for _, tc := range tests {
		got := c.Increment(tc.input)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
