package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestIntegerDivide(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a         *big.Int
		b         *big.Int
		expect    *big.Int
		expectErr bool
	}

	tests := []test{
		// Divides positive numbers.
		{a: big.NewInt(8), b: big.NewInt(2), expect: big.NewInt(4)},
		// Divides negative numbers.
		{a: big.NewInt(-8), b: big.NewInt(-2), expect: big.NewInt(4)},
		// Divides numbers in scientific notation.
		{a: big.NewInt(3e5), b: big.NewInt(2e5), expect: big.NewInt(1)},
		// Rounds positive numbers towards zero.
		{a: big.NewInt(3), b: big.NewInt(2), expect: big.NewInt(1)},
		// Rounds negative numbers towards zero.
		{a: big.NewInt(-3), b: big.NewInt(2), expect: big.NewInt(-1)},
		// Returns an error when dividing by zero.
		{a: big.NewInt(1), b: big.NewInt(0), expectErr: true},
	}

	for _, tc := range tests {
		got, err := c.IntegerDivide(tc.a, tc.b)
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
