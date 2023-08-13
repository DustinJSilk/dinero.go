package dinero

// Return true if d has a negative amount.
func (d Dinero[T]) IsNegative() bool {
	c := d.Calculator()
	return c.LessThan(d.Amount, c.Zero())
}
