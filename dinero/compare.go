package dinero

import (
	"fmt"

	"dinero.go/calculator"
)

func unsafeCompare[T any](dinero, comparator Dinero[T]) calculator.CompareResult {
	return dinero.calculator.Compare(dinero.Amount, comparator.Amount)
}

// Compare the value of d relative to comparator.
// Returns one of LT, EQ, or GT depending on whether d is less than, equal to, or greater than
// comparator.
//
// You can only compare objects that share the same currency. The function also normalizes objects
// to the same scale (the highest) before comparing them.
func (d Dinero[T]) Compare(comparator Dinero[T]) (calculator.CompareResult, error) {
	if !HaveSameCurrency(d, comparator) {
		return calculator.EQ, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, comparator)

	return unsafeCompare(normalized[0], normalized[1]), nil
}
