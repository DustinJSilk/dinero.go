package dinero

// Return true if d has a zero amount.
func (d Dinero[T]) IsZero() bool {
	c := d.calc()
	return c.Equal(d.Amount, c.Zero())
}
