package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestDecrement(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		input  *big.Int
		expect *big.Int
	}

	tests := []test{
		{input: big.NewInt(2), expect: big.NewInt(1)},
		{input: big.NewInt(-2), expect: big.NewInt(-3)},
	}

	for _, tc := range tests {
		got := c.Decrement(tc.input)

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
