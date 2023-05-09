package divide_test

import (
	"testing"

	"dinero.go/package/calculator"
	"dinero.go/package/calculator/integer"
	"dinero.go/package/divide"
)

func TestUp(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		amount    int
		factor    int
		expect    int
		expectErr bool
	}

	tests := []test{
		{amount: 20, factor: 10, expect: 2},
		{amount: -20, factor: 10, expect: -2},
		{amount: 0, factor: 10, expect: 0},
		{amount: 15, factor: 10, expect: 2},
		{amount: -15, factor: 10, expect: -1},
		{amount: 10, factor: 0, expectErr: true},
		// rounds up with any positive float quotient above half
		{amount: 1, factor: 10, expect: 1},
		{amount: 2, factor: 10, expect: 1},
		{amount: 3, factor: 10, expect: 1},
		{amount: 4, factor: 10, expect: 1},
		{amount: 5, factor: 10, expect: 1},
		{amount: 6, factor: 10, expect: 1},
		{amount: 7, factor: 10, expect: 1},
		{amount: 8, factor: 10, expect: 1},
		{amount: 9, factor: 10, expect: 1},
		// rounds up with any negative float quotient
		{amount: -1, factor: 10, expect: -0},
		{amount: -2, factor: 10, expect: -0},
		{amount: -3, factor: 10, expect: -0},
		{amount: -4, factor: 10, expect: -0},
		{amount: -5, factor: 10, expect: -0},
		{amount: -6, factor: 10, expect: -0},
		{amount: -7, factor: 10, expect: -0},
		{amount: -8, factor: 10, expect: -0},
		{amount: -9, factor: 10, expect: -0},
		// non-decimal factors
		{amount: 20, factor: 5, expect: 4},
		{amount: -20, factor: 5, expect: -4},
		{amount: 0, factor: 5, expect: 0},
		{amount: 3, factor: 2, expect: 2},
		{amount: -3, factor: 2, expect: -1},
		// rounds up with any positive float quotient
		{amount: 1, factor: 5, expect: 1},
		{amount: 2, factor: 5, expect: 1},
		{amount: 3, factor: 5, expect: 1},
		{amount: 4, factor: 5, expect: 1},
		// rounds up with any negative float quotient
		{amount: -1, factor: 5, expect: -0},
		{amount: -2, factor: 5, expect: -0},
		{amount: -3, factor: 5, expect: -0},
		{amount: -4, factor: 5, expect: -0},
	}

	for _, tc := range tests {
		got, err := divide.Up[int]{}.Divide(tc.amount, tc.factor, c)
		if err != nil {
			if tc.expectErr {
				continue
			}

			t.Fatalf("error: %v: %v, got: %v", err, tc.expect, got)
		}

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
