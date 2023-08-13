package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

func TestMultiply(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		multiplier  int
		expect      dinero.Dinero[int]
	}

	tests := []test{
		{
			description: "multiplies positive Dinero objects by positive",
			dinero:      dinero.NewDinero(400, currency.USD),
			multiplier:  4,
			expect:      dinero.NewDinero(1600, currency.USD),
		},
		{
			description: "multiplies positive Dinero objects by negative number",
			dinero:      dinero.NewDinero(400, currency.USD),
			multiplier:  -1,
			expect:      dinero.NewDinero(-400, currency.USD),
		},
		{
			description: "multiplies negative Dinero objects by positive number",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  4,
			expect:      dinero.NewDinero(-1600, currency.USD),
		},
		{
			description: "multiplies negative Dinero objects by negative number",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  -1,
			expect:      dinero.NewDinero(400, currency.USD),
		},
		{
			description: "multiplies negative Dinero objects by positive 1",
			dinero:      dinero.NewDinero(-400, currency.USD),
			multiplier:  1,
			expect:      dinero.NewDinero(-400, currency.USD),
		},
	}

	for _, tc := range tests {
		got := tc.dinero.Multiply(tc.multiplier)
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

func TestMultiplyBigInt(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[*big.Int]
		multiplier  *big.Int
		expect      dinero.Dinero[*big.Int]
	}

	tests := []test{
		{
			description: "multiplies positive Dinero objects by positive",
			dinero:      dinero.NewBigDinero(400, BigUSD),
			multiplier:  big.NewInt(4),
			expect:      dinero.NewBigDinero(1600, BigUSD),
		},
		{
			description: "multiplies positive Dinero objects by negative number",
			dinero:      dinero.NewBigDinero(400, BigUSD),
			multiplier:  big.NewInt(-1),
			expect:      dinero.NewBigDinero(-400, BigUSD),
		},
		{
			description: "multiplies negative Dinero objects by positive number",
			dinero:      dinero.NewBigDinero(-400, BigUSD),
			multiplier:  big.NewInt(4),
			expect:      dinero.NewBigDinero(-1600, BigUSD),
		},
		{
			description: "multiplies negative Dinero objects by negative number",
			dinero:      dinero.NewBigDinero(-400, BigUSD),
			multiplier:  big.NewInt(-1),
			expect:      dinero.NewBigDinero(400, BigUSD),
		},
		{
			description: "multiplies negative Dinero objects by positive 1",
			dinero:      dinero.NewBigDinero(-400, BigUSD),
			multiplier:  big.NewInt(1),
			expect:      dinero.NewBigDinero(-400, BigUSD),
		},
	}

	for _, tc := range tests {
		got := tc.dinero.Multiply(tc.multiplier)
		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestMultiplyScaledBigInt(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[*big.Int]
		multiplier  dinero.ScaledAmount[*big.Int]
		expect      dinero.Dinero[*big.Int]
		expectErr   bool
	}

	tests := []test{
		{
			description: "multiplies positive Dinero objects",
			dinero:      dinero.NewBigDinero(401, BigUSD),
			multiplier:  dinero.NewScaledAmount(big.NewInt(2001), big.NewInt(3)),
			expect:      dinero.NewBigDineroWithScale(802401, BigUSD, 5),
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
		da.Multiply(15)
	}
}
