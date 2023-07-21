package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestAdd(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a      *big.Int
		b      *big.Int
		expect *big.Int
	}

	tests := []test{
		{a: big.NewInt(2), b: big.NewInt(3), expect: big.NewInt(5)},
		{a: big.NewInt(-1), b: big.NewInt(-2), expect: big.NewInt(-3)},
		{a: big.NewInt(1e5), b: big.NewInt(2e5), expect: big.NewInt(300000)},
	}

	for _, tc := range tests {
		got := c.Add(tc.a, tc.b)

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
