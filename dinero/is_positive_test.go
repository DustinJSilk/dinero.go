package dinero_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestIsPositive(t *testing.T) {
	type test struct {
		description string
		value       dinero.Dinero[int]
		expect      bool
	}

	tests := []test{
		{
			description: "returns false when amount is less than 0",
			value:       dinero.NewDinero(-100, currency.USD),
			expect:      false,
		},
		{
			description: "returns true when amount is greater than 0",
			value:       dinero.NewDinero(100, currency.USD),
			expect:      true,
		},
		{
			description: "returns false when amount is equal to 0",
			value:       dinero.NewDinero(0, currency.USD),
			expect:      false,
		},
		{
			description: "returns false when amount is equal to -0",
			value:       dinero.NewDinero(-0, currency.USD),
			expect:      false,
		},
	}

	for _, tc := range tests {
		got := tc.value.IsPositive()

		if tc.expect != got {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkIsPositive(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		da.IsPositive()
	}
}
