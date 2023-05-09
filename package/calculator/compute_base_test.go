package calculator_test

import (
	"testing"

	"dinero.go/package/calculator"
	"dinero.go/package/calculator/integer"
)

func TestComputeBase(t *testing.T) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		input  []int
		expect int
	}

	tests := []test{
		{input: []int{100}, expect: 100},
		{input: []int{20, 12, 7}, expect: 1680},
	}

	for _, tc := range tests {
		got := c.ComputeBase(tc.input...)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}

func BenchmarkComputeBase(b *testing.B) {
	c := calculator.NewCalculator[int](integer.Calculator{})

	for i := 0; i < b.N; i++ {
		c.ComputeBase(20, 12, 7)
	}
}
