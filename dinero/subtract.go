package dinero

import (
	"fmt"
)

// Subtract the passed Dinero object from d.
func (d Dinero[T]) Subtract(subtrahend Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(d, subtrahend) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, subtrahend)
	amount := d.calculator.Subtract(normalized[0].amount, normalized[1].amount)
	return NewDineroWithOptions(amount, d.currency, normalized[0].scale, d.calculator), nil
}
