package calculator

func (c calculator[T]) Sign(value T) T {
	zero := c.Zero()

	if c.Equal(value, zero) {
		return zero
	}

	if c.LessThan(value, zero) {
		return c.Decrement(zero)
	}

	return c.One()
}
