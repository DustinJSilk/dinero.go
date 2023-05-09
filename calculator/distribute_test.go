package calculator_test

import (
	"reflect"
	"testing"

	"dinero.go/calculator"
	"dinero.go/calculator/integer"
)

func TestDistribute(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		value  int
		ratios []int
		expect []int
	}

	tests := []test{
		{value: 1003, ratios: []int{50, 50}, expect: []int{502, 501}},
		{value: 100, ratios: []int{1, 3}, expect: []int{25, 75}},
		{value: -1003, ratios: []int{50, 50}, expect: []int{-502, -501}},
		{value: 1003, ratios: []int{0, 50, 50}, expect: []int{0, 502, 501}},
		{value: 1003, ratios: []int{0, 0}, expect: []int{0, 0}},
		{value: 1003, ratios: []int{-50, -50}, expect: []int{502, 501}},
		{value: -1003, ratios: []int{-50, -50}, expect: []int{-502, -501}},
		{value: 1003, ratios: []int{}, expect: []int{}},
		{value: 0, ratios: []int{1, 10}, expect: []int{0, 0}},
	}

	for _, tc := range tests {
		got := calculator.Distribute(tc.value, tc.ratios...)

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
