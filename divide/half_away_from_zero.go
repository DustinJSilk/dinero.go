package divide

import (
	"math/big"

	"github.com/DustinJSilk/dinero.go/calculator"
)

// Divide and round towards "nearest neighbor" unless both neighbors are equidistant,
// in which case round away from zero.
type HalfAwayFromZero[T any] struct{}

var HalfAwayFromZeroInt = HalfAwayFromZero[int]{}
var HalfAwayFromZeroBigInt = HalfAwayFromZero[*big.Int]{}

func (HalfAwayFromZero[T]) Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error) {
	if !calculator.IsHalf(amount, factor) {
		return HalfUp[T]{}.Divide(amount, factor, calculator)
	}

	upAmount, err := Up[T]{}.Divide(calculator.Absolute(amount), factor, calculator)
	if err != nil {
		return amount, err
	}

	return calculator.Multiply(calculator.Sign(amount), upAmount), nil
}
