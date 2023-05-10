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
	amount := d.calculator.Add(normalized[0].amount, normalized[1].amount)
	return NewDineroWithOptions(amount, d.currency, d.scale, d.calculator), nil
}
