package dinero

func HaveSameCurrency[T any](dineros ...Dinero[T]) bool {
	c := dineros[0].calc()
	comparator := dineros[0].Currency
	comparatorBase := c.ComputeBase(comparator.Base)

	for _, v := range dineros {
		if v.Currency.Code != comparator.Code ||
			!c.Equal(v.Currency.Exponent, comparator.Exponent) ||
			!c.Equal(c.ComputeBase(v.Currency.Base), comparatorBase) {
			return false
		}
	}

	return true
}
