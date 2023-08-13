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

	c := d.Calculator()
	normalized := NormalizeScale(d, subtrahend)
	amount := c.Subtract(normalized[0].Amount, normalized[1].Amount)
	return NewDineroWithOptions(amount, d.Currency, normalized[0].Scale, c), nil
}
