package dinero

// Returns true if d is greater than or equal to comparator.
// It will always return false if they have different currencies.
func (d Dinero[T]) GreaterThanOrEqual(dinero Dinero[T]) bool {
	if !HaveSameCurrency(d, dinero) {
		return false
	}

	normalized := NormalizeScale(d, dinero)

	return d.calc().GreaterThanOrEqual(normalized[0].Amount, normalized[1].Amount)
}
