package dinero

// Return true if d has a zero amount.
func (d Dinero[T]) IsZero() bool {
	return d.Calculator.Equal(d.Amount, d.Calculator.Zero())
}
