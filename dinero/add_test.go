package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestAdd(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "adds up positive Dinero objects",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(100, currency.USD),
			expect:      dinero.NewDinero(600, currency.USD),
		},
		{
			description: "adds up negative Dinero objects",
			a:           dinero.NewDinero(-500, currency.USD),
			b:           dinero.NewDinero(-100, currency.USD),
			expect:      dinero.NewDinero(-600, currency.USD),
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDineroWithScale(1000, currency.USD, 3),
			expect:      dinero.NewDineroWithScale(6000, currency.USD, 3),
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(1000, currency.EUR),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := tc.a.Add(tc.b)
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

func BenchmarkAdd(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.Add(db)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
