package calculator_test

import (
	"reflect"
	"testing"

	"dinero.go/calculator"
	"dinero.go/calculator/integer"
)

func TestGetDivisors(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		bases  []int
		expect []int
	}

	tests := []test{
		{bases: []int{100}, expect: []int{100}},
		{bases: []int{20, 12}, expect: []int{240, 12}},
		{bases: []int{20, 12, 7}, expect: []int{1680, 84, 7}},
	}

	for _, tc := range tests {
		got := c.GetDivisors(tc.bases...)

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
