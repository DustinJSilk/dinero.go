package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestMultiply(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		multiplier  int
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(400, currency.USD),
			multiplier:  4,
			expect:      dinero.NewDinero(1600, currency.USD),
		},
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(400, currency.USD),
			multiplier:  -1,
			expect:      dinero.NewDinero(-400, currency.USD),
		},
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  4,
			expect:      dinero.NewDinero(-1600, currency.USD),
		},
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  -1,
			expect:      dinero.NewDinero(400, currency.USD),
		},
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  1,
			expect:      dinero.NewDinero(-400, currency.USD),
		},
	}

	for _, tc := range tests {
		got, err := tc.dinero.Multiply(tc.multiplier)
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

func TestMultiplyScaled(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		multiplier  dinero.ScaledAmount[int]
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewDinero(401, currency.USD),
			multiplier:  dinero.NewScaledAmount(2001, 3),
			expect:      dinero.NewDineroWithScale(802401, currency.USD, 5),
		},
	}

	for _, tc := range tests {
		got, err := tc.dinero.MultiplyScaled(tc.multiplier)
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

func BenchmarkMultiply(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.Multiply(15)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
