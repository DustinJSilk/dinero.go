package divide

import (
	"github.com/DustinJSilk/dinero.go/calculator"
)

// Divide and round towards "nearest neighbor" unless both neighbors are equidistant, in which case round down.
//
// Rounding down happens when:
// - The quotient is half (e.g., -1.5, 1.5).
// - The quotient is positive and less than half (e.g., 1.4).
// - The quotient is negative and greater than half (e.g., -1.6).
type HalfDown[T any] struct{}

var HalfDownInt = HalfDown[int]{}

func (HalfDown[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	if calculator.IsHalf(amount, factor) {
		return Down[T]{}.Divide(amount, factor, calculator)
	}

	return HalfUp[T]{}.Divide(amount, factor, calculator)
}
