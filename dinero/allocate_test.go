package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
)

func TestAllocate(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		ratios      []int
		expect      []dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		// decimal based currencies (USD)
		{
			description: "allocates to percentages",
			dinero:      dinero.NewDinero(1003, currency.USD),
			ratios:      []int{50, 50},
			expect: []dinero.Dinero[int]{
				dinero.NewDinero(502, currency.USD),
				dinero.NewDinero(501, currency.USD),
			},
		},
		{
			description: "allocates to ratios",
			dinero:      dinero.NewDinero(100, currency.USD),
			ratios:      []int{1, 3},
			expect: []dinero.Dinero[int]{
				dinero.NewDinero(25, currency.USD),
				dinero.NewDinero(75, currency.USD),
			},
		},
		{
			description: "ignores zero ratios",
			dinero:      dinero.NewDinero(1003, currency.USD),
			ratios:      []int{0, 50, 50},
			expect: []dinero.Dinero[int]{
				dinero.NewDinero(0, currency.USD),
				dinero.NewDinero(502, currency.USD),
				dinero.NewDinero(501, currency.USD),
			},
		},
		{
			description: "errors when using empty ratios",
			dinero:      dinero.NewDinero(1003, currency.USD),
			ratios:      []int{},
			expectErr:   true,
		},
		{
			description: "errors when using negative ratios",
			dinero:      dinero.NewDinero(1003, currency.USD),
			ratios:      []int{-50, 50},
			expectErr:   true,
		},
		{
			description: "errors when using only zero ratios",
			dinero:      dinero.NewDinero(1003, currency.USD),
			ratios:      []int{0, 0},
			expectErr:   true,
		},
		// non-decimal based currencies (MGA)
		{
			description: "allocates to percentages",
			dinero:      dinero.NewDinero(1003, currency.MGA),
			ratios:      []int{50, 50},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(502, currency.MGA, 1),
				dinero.NewDineroWithScale(501, currency.MGA, 1),
			},
		},
		{
			description: "allocates to ratios",
			dinero:      dinero.NewDinero(100, currency.MGA),
			ratios:      []int{1, 3},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(25, currency.MGA, 1),
				dinero.NewDineroWithScale(75, currency.MGA, 1),
			},
		},
		{
			description: "ignores zero ratios",
			dinero:      dinero.NewDinero(1003, currency.MGA),
			ratios:      []int{0, 50, 50},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(0, currency.MGA, 1),
				dinero.NewDineroWithScale(502, currency.MGA, 1),
				dinero.NewDineroWithScale(501, currency.MGA, 1),
			},
		},
		{
			description: "errors when using empty ratios",
			dinero:      dinero.NewDinero(1003, currency.MGA),
			ratios:      []int{},
			expectErr:   true,
		},
		{
			description: "errors when using negative ratios",
			dinero:      dinero.NewDinero(1003, currency.MGA),
			ratios:      []int{-50, 50},
			expectErr:   true,
		},
		{
			description: "errors when using only zero ratios",
			dinero:      dinero.NewDinero(1003, currency.MGA),
			ratios:      []int{0, 0},
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		got, err := tc.dinero.Allocate(tc.ratios...)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v", tc.description, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func TestAllocateScaled(t *testing.T) {
	type test struct {
		description string
		dinero      dinero.Dinero[int]
		ratios      []dinero.ScaledAmount[int]
		expect      []dinero.Dinero[int]
		expectErr   bool
	}

	tests := []test{
		// decimal based currencies (USD)
		{
			description: "converts the allocated amounts to the safest scale",
			dinero:      dinero.NewDinero(100, currency.USD),
			ratios: []dinero.ScaledAmount[int]{
				dinero.NewScaledAmount(505, 1),
				dinero.NewScaledAmount(495, 1),
			},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(505, currency.USD, 3),
				dinero.NewDineroWithScale(495, currency.USD, 3),
			},
		},
		{
			description: "converts the ratios to the same scale before allocating",
			dinero:      dinero.NewDinero(100, currency.USD),
			ratios: []dinero.ScaledAmount[int]{
				dinero.NewScaledAmount(5050, 2),
				dinero.NewScaledAmount(495, 1),
			},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(5050, currency.USD, 4),
				dinero.NewDineroWithScale(4950, currency.USD, 4),
			},
		},
		// non-decimal based currencies (MGA)
		{
			description: "converts the allocated amounts to the safest scale",
			dinero:      dinero.NewDinero(5, currency.MGA),
			ratios: []dinero.ScaledAmount[int]{
				dinero.NewScaledAmount(505, 1),
				dinero.NewScaledAmount(495, 1),
			},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(13, currency.MGA, 2),
				dinero.NewDineroWithScale(12, currency.MGA, 2),
			},
		},
		{
			description: "converts the ratios to the same scale before allocating",
			dinero:      dinero.NewDinero(5, currency.MGA),
			ratios: []dinero.ScaledAmount[int]{
				dinero.NewScaledAmount(5050, 2),
				dinero.NewScaledAmount(495, 1),
			},
			expect: []dinero.Dinero[int]{
				dinero.NewDineroWithScale(64, currency.MGA, 3),
				dinero.NewDineroWithScale(61, currency.MGA, 3),
			},
		},
	}

	for _, tc := range tests {
		got, err := tc.dinero.AllocateScaled(tc.ratios...)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("%s error: %v", tc.description, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("%s expected a: %v, got: %v", tc.description, tc.expect, got)
		}
	}
}

func BenchmarkAllocate(b *testing.B) {
	da := dinero.NewDinero(100, currency.USD)

	for i := 0; i < b.N; i++ {
		_, err := da.Allocate(50, 20, 10)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
