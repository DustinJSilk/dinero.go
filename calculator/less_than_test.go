package calculator_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestLessThan(t *testing.T) {
	calculator := calculator.NewCalculator[int](integer.Calculator{})

	type test struct {
		subject    int
		comparator int
		expect     bool
	}

	tests := []test{
		{subject: 4, comparator: 3, expect: false},
		{subject: -2, comparator: -3, expect: false},
		{subject: 2e5, comparator: 1e5, expect: false},
		{subject: 1, comparator: 2, expect: true},
		{subject: -3, comparator: -2, expect: true},
		{subject: 1e5, comparator: 2e5, expect: true},
	}

	for _, tc := range tests {
		got := calculator.LessThan(tc.subject, tc.comparator)

		if tc.expect != got {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}
