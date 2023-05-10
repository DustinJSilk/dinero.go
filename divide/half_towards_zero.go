package divide

import (
	"dinero.go/calculator"
)

// Divide and round towards "nearest neighbor" unless both neighbors are equidistant,
// in which case round towards zero.
type HalfTowardsZero[T any] struct{}

var HalfTowardsZeroInt = HalfTowardsZero[int]{}

func (HalfTowardsZero[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	if !calculator.IsHalf(amount, factor) {
		return HalfUp[T]{}.Divide(amount, factor, calculator)
	}

	downAmount, err := Down[T]{}.Divide(calculator.Absolute(amount), factor, calculator)
	if err != nil {
		return amount, err
	}

	return calculator.Multiply(calculator.Sign(amount), downAmount), nil
}
