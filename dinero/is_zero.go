package dinero

// Return true if d has a zero amount.
func (d Dinero[T]) IsZero() bool {
	return d.calculator.Equal(d.Amount, d.calculator.Zero())
}
