package divide

import "dinero.go/package/types"

// Divide and round up. Rounding up happens whenever the quotient is not an integer.
type Up[T any] struct{}

var UpInt = Up[int]{}

func (Up[T]) Divide(amount T, factor T, calculator types.Calculator[T]) (T, error) {
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

	if !isInteger && isPositive {
		return calculator.Increment(quotient), nil
	}

	return quotient, nil
}
