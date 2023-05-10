package dinero

import (
	"fmt"
)

// Add addend to d and return a new Dinero.
func (d Dinero[T]) Add(addend Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(d, addend) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, addend)
	amount := d.calculator.Add(normalized[0].Amount, normalized[1].Amount)
	return NewDineroWithOptions(amount, d.Currency, normalized[0].Scale, d.calculator), nil
}
