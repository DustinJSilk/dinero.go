package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestMultiply(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a      *big.Int
		b      *big.Int
		expect *big.Int
	}

	tests := []test{
		{a: big.NewInt(10), b: big.NewInt(20), expect: big.NewInt(200)},
		{a: big.NewInt(-10), b: big.NewInt(-20), expect: big.NewInt(200)},
		{a: big.NewInt(1e5), b: big.NewInt(2e5), expect: big.NewInt(20000000000)},
	}

	for _, tc := range tests {
		got := c.Multiply(tc.a, tc.b)

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
