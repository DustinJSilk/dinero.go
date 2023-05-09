package calculator_test

import (
	"testing"

	"dinero.go/package/calculator"
	"dinero.go/package/calculator/integer"
)

func TestGreaterThan(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		subject    int
		comparator int
		expect     bool
	}

	tests := []test{
		{subject: 4, comparator: 3, expect: true},
		{subject: -2, comparator: -3, expect: true},
		{subject: 2e5, comparator: 1e5, expect: true},
		{subject: 1, comparator: 2, expect: false},
		{subject: -3, comparator: -2, expect: false},
		{subject: 1e5, comparator: 2e5, expect: false},
	}

	for _, tc := range tests {
		got := calculator.GreaterThan(tc.subject, tc.comparator)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
