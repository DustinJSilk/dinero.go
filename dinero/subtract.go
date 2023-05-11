package dinero

import (
	"fmt"
)

// Subtract the passed Dinero object from d.
//
// You can only subtract objects that share the same currency. The function also normalizes objects
// to the same scale (the highest) before subtracting them.
func (d Dinero[T]) Subtract(subtrahend Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(d, subtrahend) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, subtrahend)
	amount := d.calculator.Subtract(normalized[0].Amount, normalized[1].Amount)
	return NewDineroWithOptions(amount, d.Currency, normalized[0].Scale, d.calculator), nil
}
