package dinero

// Return true if d has a positive amount.
func (d Dinero[T]) IsPositive() bool {
	return d.calculator.GreaterThan(d.amount, d.calculator.Zero())
}
