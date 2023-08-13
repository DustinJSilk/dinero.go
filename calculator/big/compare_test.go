package big_test

import (
	"math/big"
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestCompare(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a      *big.Int
		b      *big.Int
		expect calculator.CompareResult
	}

	tests := []test{
		{a: big.NewInt(1), b: big.NewInt(2), expect: calculator.LT},
		{a: big.NewInt(-2), b: big.NewInt(-1), expect: calculator.LT},
		{a: big.NewInt(1e5), b: big.NewInt(2e5), expect: calculator.LT},
		{a: big.NewInt(2), b: big.NewInt(1), expect: calculator.GT},
		{a: big.NewInt(-1), b: big.NewInt(-2), expect: calculator.GT},
		{a: big.NewInt(2e5), b: big.NewInt(1e5), expect: calculator.GT},
		{a: big.NewInt(2), b: big.NewInt(2), expect: calculator.EQ},
		{a: big.NewInt(-2), b: big.NewInt(-2), expect: calculator.EQ},
		{a: big.NewInt(2e5), b: big.NewInt(2e5), expect: calculator.EQ},
	}

	for _, tc := range tests {
		got := c.Compare(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
