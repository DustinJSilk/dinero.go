package calculator

func (c Calculator[T]) IsEven(value T) bool {
	two := c.Increment(c.One())
	remainder, _ := c.Modulo(value, two)
	return c.Equal(remainder, c.Zero())
}
