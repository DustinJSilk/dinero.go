package dinero_test

import (
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
