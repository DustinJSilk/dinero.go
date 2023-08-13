package dinero

import (
	"fmt"
)

// Add addend to d and return a new Dinero.
//
// You can only add objects that share the same currency. The function also normalizes objects to
// the same scale (the highest) before adding them up.
func (d Dinero[T]) Add(addend Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(d, addend) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	c := d.calc()
	normalized := NormalizeScale(d, addend)
	amount := c.Add(normalized[0].Amount, normalized[1].Amount)
	return NewDineroWithOptions(amount, d.Currency, normalized[0].Scale, c), nil
}
