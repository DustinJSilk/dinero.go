package dinero_test

import (
	"fmt"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestToDecimal(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[int]
		expect      string
		expectErr   bool
		transformer dinero.Transformer[int]
	}

	tests := []test{
		// Decimal currencies
		{
			description: "returns the amount in decimal format",
			value:       dinero.NewDinero(1050, currency.USD),
			expect:      "10.50",
		},
		{
			description: "returns the amount in decimal format based on a custom scale",
			value:       dinero.NewDineroWithScale(10545, currency.USD, 3),
			expect:      "10.545",
		},
		{
			description: "returns the amount in decimal format with trailing zeros",
			value:       dinero.NewDinero(1000, currency.USD),
			expect:      "10.00",
		},
		{
			description: "returns the amount in decimal format with leading zeros",
			value:       dinero.NewDinero(1005, currency.USD),
			expect:      "10.05",
		},
		{
			description: "returns the amount in decimal format and pads the decimal part",
			value:       dinero.NewDinero(500, currency.USD),
			expect:      "5.00",
		},
		{
			description: "returns the negative amount in decimal format",
			value:       dinero.NewDinero(-1050, currency.USD),
			expect:      "-10.50",
		},
		{
			description: "returns the negative amount with a leading zero in decimal format",
			value:       dinero.NewDinero(-1, currency.USD),
			expect:      "-0.01",
		},
		{
			description: "returns negative zero amount as a positive value in decimal format",
			value:       dinero.NewDinero(-0, currency.USD),
			expect:      "0.00",
		},
		{
			description: "uses a custom transformer",
			value:       dinero.NewDinero(1050, currency.USD),
			transformer: func(value string, currency currency.Currency[int]) string {
				return fmt.Sprintf("%s %s", currency.Code, value)
			},
			expect: "USD 10.50",
		},
		// Non-decimal currencies
		{
			description: "throws when passing a Dinero object using a non-decimal currency",
			value:       dinero.NewDinero(13, currency.MGA),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		opts := []dinero.Option[int]{}
		if tc.transformer != nil {
			opts = append(opts, dinero.WithTransformer(tc.transformer))
		}

		got, err := tc.value.ToDecimal(opts...)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v, %v", tc.description, tc.value, tc.expect, err)
		}

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkToDecimal(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.ToDecimal()
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
