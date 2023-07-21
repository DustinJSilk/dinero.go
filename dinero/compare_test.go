package dinero_test

import (
	"math/big"
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestCompare(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      calculator.CompareResult
		expectErr   bool
	}

	tests := []test{
		// decimal based currencies (USD)
		{
			description: "returns LT when the first amount is less than the other",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      calculator.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      calculator.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewDinero(800, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      calculator.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDineroWithScale(5000, currency.USD, 3),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      calculator.LT,
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
			expect:      calculator.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      calculator.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewDinero(8, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      calculator.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDineroWithScale(25, currency.MGA, 2),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      calculator.LT,
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

func TestCompareBigInt(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[*big.Int]
		b           dinero.Dinero[*big.Int]
		expect      calculator.CompareResult
		expectErr   bool
	}

	tests := []test{
		// decimal based currencies (USD)
		{
			description: "returns LT when the first amount is less than the other",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(800, BigUSD),
			expect:      calculator.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(500, BigUSD),
			expect:      calculator.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewBigDinero(800, BigUSD),
			b:           dinero.NewBigDinero(500, BigUSD),
			expect:      calculator.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewBigDineroWithScale(5000, BigUSD, 3),
			b:           dinero.NewBigDinero(800, BigUSD),
			expect:      calculator.LT,
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewBigDinero(800, BigUSD),
			b:           dinero.NewBigDinero(500, BigEUR),
			expectErr:   true,
		},
		// non-decimal currencies
		{
			description: "returns LT when the first amount is less than the other",
			a:           dinero.NewBigDinero(5, BigMGA),
			b:           dinero.NewBigDinero(8, BigMGA),
			expect:      calculator.LT,
		},
		{
			description: "returns EQ when amounts are equal",
			a:           dinero.NewBigDinero(5, BigMGA),
			b:           dinero.NewBigDinero(5, BigMGA),
			expect:      calculator.EQ,
		},
		{
			description: "returns GT when the first amount is greater than the other",
			a:           dinero.NewBigDinero(8, BigMGA),
			b:           dinero.NewBigDinero(5, BigMGA),
			expect:      calculator.GT,
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewBigDineroWithScale(25, BigMGA, 2),
			b:           dinero.NewBigDinero(8, BigMGA),
			expect:      calculator.LT,
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewBigDinero(800, BigUSD),
			b:           dinero.NewBigDinero(5, BigMGA),
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
