package dinero

// Returns true if d is less than comparator.
// It will always return false if they have different currencies.
func (d Dinero[T]) LessThan(dinero Dinero[T]) bool {
	if !HaveSameCurrency(d, dinero) {
		return false
	}

	normalized := NormalizeScale(d, dinero)

	return d.Calculator().LessThan(normalized[0].Amount, normalized[1].Amount)
}
