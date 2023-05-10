package dinero

// Returns true if d is less than or equal to comparator.
// It will always return false if they have different currencies.
func (d Dinero[T]) LessThanOrEqual(dinero Dinero[T]) bool {
	if !HaveSameCurrency(d, dinero) {
		return false
	}

	normalized := NormalizeScale(d, dinero)

	return d.calculator.LessThanOrEqual(normalized[0].amount, normalized[1].amount)
}
