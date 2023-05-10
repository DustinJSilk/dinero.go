package dinero

// Check whether the value of d is equal to comparator.
func (d Dinero[T]) Equal(comparator Dinero[T]) bool {
	return HaveSameAmount(d, comparator) && HaveSameCurrency(d, comparator)
}
