package divide

import (
	"dinero.go/calculator"
)

// Divide and round towards "nearest neighbor" unless both neighbors are
// equidistant, in which case round up.
//
// Rounding up happens when:
// - The quotient is half (e.g., -1.5, 1.5).
// - The quotient is positive and greater than half (e.g., 1.6).
// - The quotient is negative and less than half (e.g., -1.4).
type HalfUp[T any] struct{}

var HalfUpInt = HalfUp[int]{}

func (HalfUp[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	zero := calculator.Zero()
	remainder, err := calculator.Modulo(amount, factor)
	if err != nil {
		return amount, err
	}

	absoluteRemainder := calculator.Absolute(remainder)
	difference := calculator.Subtract(factor, absoluteRemainder)
	isLessThanHalf := calculator.GreaterThan(difference, absoluteRemainder)
	isPositive := calculator.GreaterThan(amount, zero)

	if calculator.IsHalf(amount, factor) ||
		(isPositive && !isLessThanHalf) ||
		(!isPositive && isLessThanHalf) {
		return Up[T]{}.Divide(amount, factor, calculator)
	}

	return Down[T]{}.Divide(amount, factor, calculator)
}
