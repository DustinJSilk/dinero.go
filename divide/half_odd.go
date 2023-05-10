package divide

import (
	"dinero.go/calculator"
)

// Divide and round towards "nearest neighbor" unless both neighbors are
// equidistant, in which case round to the nearest odd integer.
type HalfOdd[T any] struct{}

var HalfOddInt = HalfOdd[int]{}

func (HalfOdd[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	rounded, err := HalfUp[T]{}.Divide(amount, factor, calculator)
	if err != nil {
		return amount, err
	}

	if !calculator.IsHalf(amount, factor) {
		return rounded, nil
	} else if calculator.IsEven(rounded) {
		return calculator.Decrement(rounded), nil
	}

	return rounded, nil
}
