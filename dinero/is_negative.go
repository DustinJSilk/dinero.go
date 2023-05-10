package dinero

// Return true if d has a negative amount.
func (d Dinero[T]) IsNegative() bool {
	return d.calculator.LessThan(d.amount, d.calculator.Zero())
}
