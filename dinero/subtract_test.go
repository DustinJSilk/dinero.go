package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestSubtract(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[int]
		b           dinero.Dinero[int]
		expect      dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "subtracts positive Dinero objects",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(100, currency.USD),
			expect:      dinero.NewDinero(400, currency.USD),
		},
		{
			description: "subtracts negative Dinero objects",
			a:           dinero.NewDinero(-500, currency.USD),
			b:           dinero.NewDinero(-100, currency.USD),
			expect:      dinero.NewDinero(-400, currency.USD),
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDineroWithScale(1000, currency.USD, 3),
			expect:      dinero.NewDineroWithScale(4000, currency.USD, 3),
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewDinero(500, currency.USD),
			b:           dinero.NewDinero(1000, currency.EUR),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := tc.a.Subtract(tc.b)
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

func TestSubtractBigInt(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[*big.Int]
		b           dinero.Dinero[*big.Int]
		expect      dinero.Dinero[*big.Int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "subtracts positive Dinero objects",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(100, BigUSD),
			expect:      dinero.NewBigDinero(400, BigUSD),
		},
		{
			description: "subtracts negative Dinero objects",
			a:           dinero.NewBigDinero(-500, BigUSD),
			b:           dinero.NewBigDinero(-100, BigUSD),
			expect:      dinero.NewBigDinero(-400, BigUSD),
		},
		{
			description: "normalizes the result to the highest scale",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDineroWithScale(1000, BigUSD, 3),
			expect:      dinero.NewBigDineroWithScale(4000, BigUSD, 3),
		},
		{
			description: "errors when using different currencies",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(1000, BigEUR),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := tc.a.Subtract(tc.b)
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

func BenchmarkSubtract(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.Subtract(db)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
