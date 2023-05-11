package dinero

// Check whether the value of a Dinero object is equal to another.
//
// This function does same-value equality, determining whether two Dinero objects are functionally
// equivalent. It also normalizes objects to the same scale (the highest) before comparing them.
func (d Dinero[T]) Equal(comparator Dinero[T]) bool {
	return HaveSameAmount(d, comparator) && HaveSameCurrency(d, comparator)
}
