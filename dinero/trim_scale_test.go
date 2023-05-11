package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestTrimScale(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		multiplier  int
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "trims a Dinero object down to its currency exponent's scale",
			dinero:      dinero.NewDineroWithScale(500000, currency.USD, 5),
			expect:      dinero.NewDineroWithScale(500, currency.USD, 2),
		},
		{
			description: "trims a Dinero object down to the safest possible scale",
			dinero:      dinero.NewDineroWithScale(55550, currency.USD, 4),
			expect:      dinero.NewDineroWithScale(5555, currency.USD, 3),
		},
		{
			description: "doesn't trim the scale when there's nothing to trim",
			dinero:      dinero.NewDineroWithScale(5555, currency.USD, 3),
			expect:      dinero.NewDineroWithScale(5555, currency.USD, 3),
		},
		{
			description: "doesn't crash on zero amounts",
			dinero:      dinero.NewDinero(0, currency.USD),
			expect:      dinero.NewDineroWithScale(0, currency.USD, 2),
		},
	}

	for _, tc := range tests {
		got, err := tc.dinero.TrimScale()
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v, %v", tc.description, tc.dinero, tc.multiplier, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkTrimScale(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.TrimScale()
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
