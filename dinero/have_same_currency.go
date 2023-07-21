package dinero

func HaveSameCurrency[T any](dineros ...Dinero[T]) bool {
	calculator := dineros[0].Calculator
	comparator := dineros[0].Currency
	comparatorBase := calculator.ComputeBase(comparator.Base)

	for _, v := range dineros {
		if v.Currency.Code != comparator.Code ||
			!calculator.Equal(v.Currency.Exponent, comparator.Exponent) ||
			!calculator.Equal(calculator.ComputeBase(v.Currency.Base), comparatorBase) {
			return false
		}
	}

	return true
}
