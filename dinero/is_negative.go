package dinero

// Return true if d has a negative amount.
func (d Dinero[T]) IsNegative() bool {
	c := d.calc()
	return c.LessThan(d.Amount, c.Zero())
}
