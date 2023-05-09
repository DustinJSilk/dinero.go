package calculator

func (c Calculator[T]) Absolute(value T) T {
	zero := c.Zero()

	if c.Equal(value, zero) {
		return zero
	}

	if c.LessThan(value, zero) {
		minusOne := c.Decrement(zero)
		return c.Multiply(value, minusOne)
	}

	return value
}
