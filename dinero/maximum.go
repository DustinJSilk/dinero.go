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
	c := d.calc()
	amounts := make([]T, len(dineros))
	for i, v := range normalized {
		amounts[i] = v.Amount
	}

	return NewDineroWithOptions(
		c.Maximum(amounts...),
		d.Currency,
		d.Scale,
		c,
	), nil
}
