package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestToUnit(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[int]
		expect      []int
		expectErr   bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns the amount in currency units",
			value:       dinero.NewDinero(1050, currency.USD),
			expect:      []int{10, 50},
		},
		{
			description: "returns the amount in currency units, based on a custom scale",
			value:       dinero.NewDineroWithScale(10545, currency.USD, 3),
			expect:      []int{10, 545},
		},
		{
			description: "returns the amount in currency units, with a single trailing zero",
			value:       dinero.NewDinero(1000, currency.USD),
			expect:      []int{10, 0},
		},
		// non-decimal currencies
		{
			description: "returns the amount in currency units",
			value:       dinero.NewDinero(9, currency.NewCurrency("GRD", 6, 1)),
			expect:      []int{1, 3},
		},
	}

	for _, tc := range tests {
		got, err := tc.value.ToUnit()
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v, %v", tc.description, tc.value, tc.expect, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestToUnitBigInt(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[*big.Int]
		expect      []*big.Int
		expectErr   bool
	}

	tests := []test{
		// decimal currencies
		{
			description: "returns the amount in currency units",
			value:       dinero.NewBigDinero(1050, BigUSD),
			expect:      []*big.Int{big.NewInt(10), big.NewInt(50)},
		},
		{
			description: "returns the amount in currency units, based on a custom scale",
			value:       dinero.NewBigDineroWithScale(10545, BigUSD, 3),
			expect:      []*big.Int{big.NewInt(10), big.NewInt(545)},
		},
		{
			description: "returns the amount in currency units, with a single trailing zero",
			value:       dinero.NewBigDinero(1000, BigUSD),
			expect:      []*big.Int{big.NewInt(10), big.NewInt(0)},
		},
		// non-decimal currencies
		{
			description: "returns the amount in currency units",
			value:       dinero.NewBigDinero(9, currency.NewCurrency("GRD", big.NewInt(6), big.NewInt(1))),
			expect:      []*big.Int{big.NewInt(1), big.NewInt(3)},
		},
	}

	for _, tc := range tests {
		got, err := tc.value.ToUnit()
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v, %v", tc.description, tc.value, tc.expect, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkToUnit(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.ToUnit()
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
