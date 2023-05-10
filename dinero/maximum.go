package dinero

import (
	"fmt"
)

// Returns the greatest of the passed Dineros.
func Maximum[T any](dineros ...Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(dineros...) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(dineros...)
	d := normalized[0]
	amounts := make([]T, len(dineros))
	for i, v := range normalized {
		amounts[i] = v.amount
	}

	return NewDineroWithOptions(
		d.calculator.Maximum(amounts...),
		d.currency,
		d.scale,
		d.calculator,
	), nil
}
