package calculator

func (c calculator[T]) IsEven(value T) bool {
	two := c.Increment(c.One())
	remainder, _ := c.Modulo(value, two)
	return c.Equal(remainder, c.Zero())
}
