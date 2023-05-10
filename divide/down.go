package divide

import (
	"dinero.go/calculator"
)

// Divide and round down. Rounding down happens whenever the quotient is not an integer.
type Down[T any] struct{}

var DownInt = Down[int]{}

func (Down[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	zero := calculator.Zero()
	isPositive := calculator.GreaterThan(amount, zero)
	quotient, err := calculator.IntegerDivide(amount, factor)
	if err != nil {
		return calculator.Zero(), err
	}

	remainder, err := calculator.Modulo(amount, factor)
	if err != nil {
		return calculator.Zero(), err
	}

	isInteger := calculator.Equal(remainder, zero)

	if isPositive || isInteger {
		return quotient, nil
	}

	return calculator.Decrement(quotient), nil
}
