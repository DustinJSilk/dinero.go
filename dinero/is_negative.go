package dinero

// Return true if d has a negative amount.
func (d Dinero[T]) IsNegative() bool {
	return d.Calculator.LessThan(d.Amount, d.Calculator.Zero())
}
