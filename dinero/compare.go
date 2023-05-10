package dinero

import (
	"fmt"

	"dinero.go/types"
)

func unsafeCompare[T any](dinero, comparator Dinero[T]) types.CompareResult {
	return dinero.calculator.Compare(dinero.amount, comparator.amount)
}

// Compare the value of d relative to comparator.
// Returns one of LT, EQ, or GT depending on whether d is less than, equal to, or greater than comparator.
func (d Dinero[T]) Compare(comparator Dinero[T]) (types.CompareResult, error) {
	if !HaveSameCurrency(d, comparator) {
		return types.EQ, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, comparator)

	return unsafeCompare(normalized[0], normalized[1]), nil
}
