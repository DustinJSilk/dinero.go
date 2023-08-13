package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestConvert(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[int]
		currency    currency.Currency[int]
		rates       map[string]dinero.ScaledAmount[int]
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "converts a Dinero object to another currency",
			value:       dinero.NewDinero(500, currency.USD),
			currency:    currency.EUR,
			rates: map[string]dinero.ScaledAmount[int]{
				"EUR": {
					Amount: 89,
					Scale:  2,
				},
			},
			expect: dinero.NewDineroWithScale(44500, currency.EUR, 4),
		},
		{
			description: "uses the destination currency's exponent as scale",
			value:       dinero.NewDinero(500, currency.USD),
			currency:    currency.IQD,
			rates: map[string]dinero.ScaledAmount[int]{
				"IQD": {
					Amount: 1199,
				},
			},
			expect: dinero.NewDineroWithScale(5995000, currency.IQD, 3),
		},
		// non-decimal currencies
		{
			description: "converts a Dinero object to another currency",
			value:       dinero.NewDinero(1, currency.MRU),
			currency:    currency.MGA,
			rates: map[string]dinero.ScaledAmount[int]{
				"MGA": {
					Amount: 108,
				},
			},
			expect: dinero.NewDineroWithScale(108, currency.MGA, 1),
		},
		{
			description: "converts to the safest scale",
			value:       dinero.NewDinero(100, currency.USD),
			currency:    currency.MGA,
			rates: map[string]dinero.ScaledAmount[int]{
				"MGA": {
					Amount: 3912566,
					Scale:  3,
				},
			},
			expect: dinero.NewDineroWithScale(391256600, currency.MGA, 5),
		},
	}

	for _, tc := range tests {
		got, err := tc.value.Convert(tc.currency, tc.rates)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v", tc.description, tc.value, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestConvertBigInt(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[*big.Int]
		currency    currency.Currency[*big.Int]
		rates       map[string]dinero.ScaledAmount[*big.Int]
		expect      dinero.Dinero[*big.Int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "converts a Dinero object to another currency",
			value:       dinero.NewBigDinero(500, BigUSD),
			currency:    BigEUR,
			rates: map[string]dinero.ScaledAmount[*big.Int]{
				"EUR": {
					Amount: big.NewInt(89),
					Scale:  big.NewInt(2),
				},
			},
			expect: dinero.NewBigDineroWithScale(44500, BigEUR, 4),
		},
		{
			description: "uses the destination currency's exponent as scale",
			value:       dinero.NewBigDinero(500, BigUSD),
			currency:    BigIQD,
			rates: map[string]dinero.ScaledAmount[*big.Int]{
				"IQD": {
					Amount: big.NewInt(1199),
					Scale:  big.NewInt(0),
				},
			},
			expect: dinero.NewBigDineroWithScale(5995000, BigIQD, 3),
		},
		// non-decimal currencies
		{
			description: "converts a Dinero object to another currency",
			value:       dinero.NewBigDinero(1, BigMRU),
			currency:    BigMGA,
			rates: map[string]dinero.ScaledAmount[*big.Int]{
				"MGA": {
					Amount: big.NewInt(108),
					Scale:  big.NewInt(0),
				},
			},
			expect: dinero.NewBigDineroWithScale(108, BigMGA, 1),
		},
		{
			description: "converts to the safest scale",
			value:       dinero.NewBigDinero(100, BigUSD),
			currency:    BigMGA,
			rates: map[string]dinero.ScaledAmount[*big.Int]{
				"MGA": {
					Amount: big.NewInt(3912566),
					Scale:  big.NewInt(3),
				},
			},
			expect: dinero.NewBigDineroWithScale(391256600, BigMGA, 5),
		},
	}

	for _, tc := range tests {
		got, err := tc.value.Convert(tc.currency, tc.rates)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v", tc.description, tc.value, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	rates := map[string]dinero.ScaledAmount[int]{
		"EUR": {
			Amount: 89,
			Scale:  2,
		},
	}

	for i := 0; i < b.N; i++ {
		_, err := da.Convert(currency.EUR, rates)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
