package calculator

// Returns true if value is half of total.
func (c calculator[T]) IsHalf(value T, total T) bool {
	zero := c.Zero()
	if c.Equal(zero, value) || c.Equal(zero, total) {
		return false
	}

	remainder, _ := c.Modulo(value, total)
	absoluteRemainder := c.Absolute(remainder)
	diff := c.Subtract(total, absoluteRemainder)

	return c.Equal(diff, absoluteRemainder)
}
