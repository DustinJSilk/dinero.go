package dinero_test

import (
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestGreaterThan(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns false when the first amount is less than the other",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      false,
		},
		{
			description: "returns false when amounts are equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      false,
		},
		{
			description: "returns true when the first amount is greater than the other",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      true,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDineroWithScale(5000, currency.USD, 3),
			expect:      true,
		},
		{
			description: "return false when using different currencies",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(5000, currency.EUR),
			expect:      false,
		},
		// non-decimal currencies
		{
			description: "returns false when the first amount is less than the other",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      false,
		},
		{
			description: "returns false when amounts are equal",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      false,
		},
		{
			description: "returns true when the first amount is greater than the other",
			a:           dinero.NewDinero(8, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      true,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDinero(8, currency.MGA),
			b:           dinero.NewDineroWithScale(25, currency.MGA, 2),
			expect:      true,
		},
		{
			description: "return false when using different currencies",
			a:           dinero.NewDinero(800, currency.MGA),
			b:           dinero.NewDinero(5000, currency.EUR),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := tc.a.GreaterThan(tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkGreaterThan(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		da.GreaterThan(db)
	}
}
