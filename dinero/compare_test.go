package dinero_test

import (
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
	"dinero.go/types"
)

func TestCompare(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      types.CompareResult
		expectErr   bool
	}

	tests := []test{
		// decimal based currencies (USD)
		{
			description: "returns LT when the first amount is less than the other",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      types.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      types.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      types.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDineroWithScale(5000, currency.USD, 3),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      types.LT,
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(500, currency.EUR),
			expectErr:   true,
		},
		// non-decimal currencies
		{
			description: "returns LT when the first amount is less than the other",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      types.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      types.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewDinero(8, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      types.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDineroWithScale(25, currency.MGA, 2),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      types.LT,
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(5, currency.MGA),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := tc.a.Compare(tc.b)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v", tc.description, err)
		}

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(200, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.Compare(db)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
