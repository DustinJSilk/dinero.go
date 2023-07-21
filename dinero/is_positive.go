package dinero

// Return true if d has a positive amount.
func (d Dinero[T]) IsPositive() bool {
	return d.Calculator.GreaterThan(d.Amount, d.Calculator.Zero())
}
