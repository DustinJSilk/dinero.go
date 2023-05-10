package calculator

import (
	"dinero.go/types"
)

type Calculator[T any] struct {
	core types.CalculatorCore[T]
	one  T
	ten  T
}

func NewCalculator[T any](core types.CalculatorCore[T]) Calculator[T] {
	one := core.Increment(core.Zero())
	ten := core.Zero()
	for i := 0; i < 10; i++ {
		ten = core.Increment(ten)
	}

	return Calculator[T]{core, one, ten}
}

func (c Calculator[T]) Add(augend, addend T) T {
	return c.core.Add(augend, addend)
}

func (c Calculator[T]) Compare(a, b T) types.CompareResult {
	return c.core.Compare(a, b)
}

func (c Calculator[T]) Decrement(value T) T {
	return c.core.Decrement(value)
}

func (c Calculator[T]) Increment(value T) T {
	return c.core.Increment(value)
}

func (c Calculator[T]) IntegerDivide(dividend, divisor T) (T, error) {
	return c.core.IntegerDivide(dividend, divisor)
}

func (c Calculator[T]) Modulo(dividend, divisor T) (T, error) {
	return c.core.Modulo(dividend, divisor)
}

func (c Calculator[T]) Multiply(multiplicand, multiplier T) T {
	return c.core.Multiply(multiplicand, multiplier)
}

func (c Calculator[T]) One() T {
	return c.one
}

func (c Calculator[T]) Power(base, exponent T) T {
	return c.core.Power(base, exponent)
}

func (c Calculator[T]) Subtract(minuend, subtrahend T) T {
	return c.core.Subtract(minuend, subtrahend)
}

func (c Calculator[T]) Ten() T {
	return c.ten
}

func (c Calculator[T]) Zero() T {
	return c.core.Zero()
}
