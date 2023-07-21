package big_test

import (
	"math"
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestPower(t *testing.T) {
	c := bigcalc.Calculator{}

	type test struct {
		a      *big.Int
		b      *big.Int
		expect *big.Int
	}

	tests := []test{
		{a: big.NewInt(5), b: big.NewInt(1), expect: big.NewInt(5)},
		{a: big.NewInt(5), b: big.NewInt(0), expect: big.NewInt(1)},
		{a: big.NewInt(0), b: big.NewInt(5), expect: big.NewInt(0)},
		{a: big.NewInt(2), b: big.NewInt(3), expect: big.NewInt(8)},
		{a: big.NewInt(2), b: big.NewInt(2), expect: big.NewInt(4)},
		{a: big.NewInt(-2), b: big.NewInt(3), expect: big.NewInt(-8)},
		{a: big.NewInt(123), b: big.NewInt(7), expect: big.NewInt(425927596977747)},
		{a: big.NewInt(1e5), b: big.NewInt(3), expect: big.NewInt(1000000000000000)},
	}

	for _, tc := range tests {
		got := c.Power(tc.a, tc.b)

		if tc.expect.Cmp(got) != 0 {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}

func FuzzPower(f *testing.F) {
	c := bigcalc.Calculator{}
	f.Add(5, 3)

	f.Fuzz(func(t *testing.T, base int, exponent int) {
		powered := c.Power(big.NewInt(int64(base)), big.NewInt(int64(exponent)))

		ex := int(math.Log(float64(powered.Int64())) / math.Log(float64(base)))
		bs := int(math.Pow(float64(powered.Int64()), 1/float64(exponent)))

		if ex != exponent {
			t.Errorf("Exponent before: %q, after: %q", exponent, ex)
		}
		if bs != base {
			t.Errorf("Base before: %q, after: %q", exponent, ex)
		}
	})
}
