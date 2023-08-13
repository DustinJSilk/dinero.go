package dinero_test

import (
	"math/big"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestHaveSameAmount(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(500, currency.USD),
			expect:      true,
		},
		{
			description: "returns true when amounts are equal",
			a:           dinero.NewDinero(1000, currency.USD),
			b:           dinero.NewDinero(1000, currency.USD),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewDinero(1000, currency.USD),
			b:           dinero.NewDinero(2000, currency.USD),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal once normalized",
			a:           dinero.NewDinero(1000, currency.USD),
			b:           dinero.NewDineroWithScale(10000, currency.USD, 3),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal once normalized",
			a:           dinero.NewDinero(10000, currency.USD),
			b:           dinero.NewDineroWithScale(10000, currency.USD, 3),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := dinero.HaveSameAmount(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestHaveSameAmountBigInt(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[*big.Int]
		b           dinero.Dinero[*big.Int]
		expect      bool
	}

	tests := []test{
		{
			description: "returns true when amounts and currencies are equal",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(500, BigUSD),
			expect:      true,
		},
		{
			description: "returns true when amounts are equal",
			a:           dinero.NewBigDinero(1000, BigUSD),
			b:           dinero.NewBigDinero(1000, BigUSD),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal",
			a:           dinero.NewBigDinero(1000, BigUSD),
			b:           dinero.NewBigDinero(2000, BigUSD),
			expect:      false,
		},
		{
			description: "returns true when amounts are equal once normalized",
			a:           dinero.NewBigDinero(1000, BigUSD),
			b:           dinero.NewBigDineroWithScale(10000, BigUSD, 3),
			expect:      true,
		},
		{
			description: "returns false when amounts are not equal once normalized",
			a:           dinero.NewBigDinero(10000, BigUSD),
			b:           dinero.NewBigDineroWithScale(10000, BigUSD, 3),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := dinero.HaveSameAmount(tc.a, tc.b)

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkHaveSameAmount(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		dinero.HaveSameAmount(da, db)
	}
}
