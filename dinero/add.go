package dinero

import (
	"fmt"
)

func unsafeAdd[T any](augend, addend Dinero[T]) Dinero[T] {
	amount := augend.calculator.Add(augend.amount, addend.amount)
	return NewDineroWithOptions(amount, augend.currency, augend.scale, augend.calculator)
}

// Add addend to d and return a new Dinero.
func (d Dinero[T]) Add(addend Dinero[T]) (Dinero[T], error) {
	if !HaveSameCurrency(d, addend) {
		return Dinero[T]{}, fmt.Errorf("mismatched currencies")
	}

	normalized := NormalizeScale(d, addend)

	return unsafeAdd(normalized[0], normalized[1]), nil
}
