package dinero

import (
	"fmt"
)

// Returns the lowest of the passed Dineros.
func Minimum[T any](dineros ...Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(dineros...) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(dineros...)
	d := normalized[0]
	amounts := make([]T, len(dineros))
	for i, v := range normalized {
		amounts[i] = v.Amount
	}

	return NewDineroWithOptions(
		d.calculator.Minimum(amounts...),
		d.Currency,
		d.Scale,
		d.calculator,
	), nil
}
