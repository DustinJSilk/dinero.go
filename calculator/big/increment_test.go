package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestIncrement(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		input  *big.Int
		expect *big.Int
	}

	tests := []test{
		{input: big.NewInt(2), expect: big.NewInt(3)},
		{input: big.NewInt(-2), expect: big.NewInt(-1)},
	}

	for _, tc := range tests {
		got := c.Increment(tc.input)

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
