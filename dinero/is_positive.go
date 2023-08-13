package dinero

// Return true if d has a positive amount.
func (d Dinero[T]) IsPositive() bool {
	c := d.Calculator()
	return c.GreaterThan(d.Amount, c.Zero())
}
