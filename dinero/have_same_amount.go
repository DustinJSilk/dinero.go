package dinero

func HaveSameAmount[T any](dineros ...Dinero[T]) bool {
	normalized := NormalizeScale(dineros...)
	first := normalized[0].Amount
	c := dineros[0].calculator

	for _, v := range normalized {
		if !c.Equal(first, v.Amount) {
			return false
		}
	}

	return true
}
