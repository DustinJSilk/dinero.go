package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestMaximum(t *testing.T) {
	type test struct {
		description string
		a           dinero.Dinero[*big.Int]
		b           dinero.Dinero[*big.Int]
		expect      dinero.Dinero[*big.Int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "returns the greatest from a set of Dinero objects",
			a:           dinero.NewBigDinero(150, BigUSD),
			b:           dinero.NewBigDinero(50, BigUSD),
			expect:      dinero.NewBigDinero(150, BigUSD),
		},
		{
			description: "returns the greatest from a set of Dinero objects after normalization",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDineroWithScale(1000, BigUSD, 3),
			expect:      dinero.NewBigDineroWithScale(5000, BigUSD, 3),
		},
		{
			description: "returns an error when using different currencies",
			a:           dinero.NewBigDinero(500, BigUSD),
			b:           dinero.NewBigDinero(100, BigEUR),
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := dinero.Maximum(tc.a, tc.b)
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

func BenchmarkMaximum(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)
	db := dinero.NewDinero(150, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := dinero.Maximum(da, db)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
