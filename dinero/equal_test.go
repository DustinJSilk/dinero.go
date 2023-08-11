package dinero_test

import (
	"math/big"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestEqual(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(800, currency.USD),
			expect:      false,
		},
		{
			description: "returns false when currencies are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.EUR),
			expect:      false,
		},
		{
			description: "returns false when amounts and currencies are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(800, currency.EUR),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal after normalization",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDineroWithScale(5000, currency.USD, 3),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal after normalization",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDineroWithScale(500, currency.USD, 3),
			expect:      false,
		},
		// non-decimal currencies
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(5, currency.MGA),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      false,
		},
		{
			description: "returns false when currencies are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.MGA),
			expect:      false,
		},
		{
			description: "returns false when amounts and currencies are not equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(8, currency.MGA),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal after normalization",
			a:           dinero.NewDinero(5, currency.MGA),
			b:           dinero.NewDineroWithScale(25, currency.MGA, 2),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal after normalization",
			a:           dinero.NewDinero(25, currency.MGA),
			b:           dinero.NewDineroWithScale(25, currency.MGA, 2),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := tc.a.Equal(tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestEqualBigInt(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[*big.Int]
		b           dinero.Dinero[*big.Int]
		expect      bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(500, BigUSD),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(800, BigUSD),
			expect:      false,
		},
		{
			description: "returns false when currencies are not equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(500, BigEUR),
			expect:      false,
		},
		{
			description: "returns false when amounts and currencies are not equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(800, BigEUR),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal after normalization",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDineroWithScale(5000, BigUSD, 3),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal after normalization",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDineroWithScale(500, BigUSD, 3),
			expect:      false,
		},
		// non-decimal currencies
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewBigDinero(5, BigMGA),
			b:           dinero.NewBigDinero(5, BigMGA),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewBigDinero(5, BigMGA),
			b:           dinero.NewBigDinero(8, BigMGA),
			expect:      false,
		},
		{
			description: "returns false when currencies are not equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(500, BigMGA),
			expect:      false,
		},
		{
			description: "returns false when amounts and currencies are not equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(8, BigMGA),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal after normalization",
			a:           dinero.NewBigDinero(5, BigMGA),
			b:           dinero.NewBigDineroWithScale(25, BigMGA, 2),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal after normalization",
			a:           dinero.NewBigDinero(25, BigMGA),
			b:           dinero.NewBigDineroWithScale(25, BigMGA, 2),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := tc.a.Equal(tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkEqual(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		da.Equal(db)
	}
}
