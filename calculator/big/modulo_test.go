package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestModulo(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a         *big.Int
		b         *big.Int
		expect    *big.Int
		expectErr bool
	}

	tests := []test{
		{a: big.NewInt(5), b: big.NewInt(3), expect: big.NewInt(2)},
		{a: big.NewInt(-5), b: big.NewInt(-4), expect: big.NewInt(-1)},
		{a: big.NewInt(4e5), b: big.NewInt(3e5), expect: big.NewInt(100000)},
		{a: big.NewInt(1), b: big.NewInt(0), expectErr: true},
	}

	for _, tc := range tests {
		got, err := c.Modulo(tc.a, tc.b)
		if err != nil {
			if tc.expectErr {
				continue
			}
			t.Fatalf("error: %v, got: %v", tc.expect, got)
		}

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
