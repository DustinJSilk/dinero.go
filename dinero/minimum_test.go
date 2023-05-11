package dinero_test

import (
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestMinimum(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "returns the lowest from a set of Dinero objects",
			a:           dinero.NewDinero(150, currency.USD),
			b:           dinero.NewDinero(50, currency.USD),
			expect:      dinero.NewDinero(50, currency.USD),
		},
		{
			description: "returns the lowest from a set of Dinero objects after normalization",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDineroWithScale(1000, currency.USD, 3),
			expect:      dinero.NewDineroWithScale(1000, currency.USD, 3),
		},
		{
			description: "returns an error when using different currencies",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(100, currency.EUR),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := dinero.Minimum(tc.a, tc.b)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v, %v, %v", tc.description, tc.a, tc.b, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkMinimum(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(150, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := dinero.Minimum(da, db)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
