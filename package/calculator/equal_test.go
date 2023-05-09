package calculator_test

import (
	"testing"

	"dinero.go/package/calculator"
	"dinero.go/package/calculator/integer"
)

func TestEqual(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		subject    int
		comparator int
		expect     bool
	}

	tests := []test{
		{subject: 2, comparator: 2, expect: true},
		{subject: -2, comparator: -2, expect: true},
		{subject: 1e5, comparator: 1e5, expect: true},
		{subject: 2, comparator: 3, expect: false},
		{subject: -2, comparator: -3, expect: false},
		{subject: 1e5, comparator: 2e5, expect: false},
	}

	for _, tc := range tests {
		got := calculator.Equal(tc.subject, tc.comparator)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
