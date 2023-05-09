package calculator_test

import (
	"testing"

	"dinero.go/package/calculator"
	"dinero.go/package/calculator/integer"
)

func TestLessThanOrEqual(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		subject    int
		comparator int
		expect     bool
	}

	tests := []test{
		{subject: 2, comparator: 3, expect: true},
		{subject: -4, comparator: -3, expect: true},
		{subject: 2e5, comparator: 4e5, expect: true},
		{subject: 2, comparator: 2, expect: true},
		{subject: -2, comparator: -2, expect: true},
		{subject: 2e5, comparator: 2e5, expect: true},
		{subject: 3, comparator: 2, expect: false},
		{subject: -3, comparator: -4, expect: false},
		{subject: 3e5, comparator: 2e5, expect: false},
	}

	for _, tc := range tests {
		got := calculator.LessThanOrEqual(tc.subject, tc.comparator)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
