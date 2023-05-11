package integer_test

import (
	"math"
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestPower(t *testing.T) {
	c := integer.Calculator{}

	type test struct {
		a      int
		b      int
		expect int
	}

	tests := []test{
		{a: 5, b: 1, expect: 5},
		{a: 5, b: 0, expect: 1},
		{a: 0, b: 5, expect: 0},
		{a: 2, b: 3, expect: 8},
		{a: 2, b: 2, expect: 4},
		{a: -2, b: 3, expect: -8},
		{a: 123, b: 7, expect: 425927596977747},
		{a: 1e5, b: 3, expect: 1000000000000000},
	}

	for _, tc := range tests {
		got := c.Power(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}

func FuzzPower(f *testing.F) {
	c := integer.Calculator{}
	f.Add(5, 3)

	f.Fuzz(func(t *testing.T, base int, exponent int) {
		powered := c.Power(base, exponent)

		ex := int(math.Log(float64(powered)) / math.Log(float64(base)))
		bs := int(math.Pow(float64(powered), 1/float64(exponent)))

		if ex != exponent {
			t.Errorf("Exponent before: %q, after: %q", exponent, ex)
		}
		if bs != base {
			t.Errorf("Base before: %q, after: %q", exponent, ex)
		}
	})
}
