package dinero

func HaveSameAmount[T any](dineros ...Dinero[T]) bool {
	// normalized := NormalizeScale(dineros...)

	// calculator := dineros[0].calculator
	// comparator := dineros[0].currency
	// comparatorBase := calculator.ComputeBase(comparator.Base())

	// for _, v := range dineros {
	// 	if v.currency.Code() != comparator.Code() ||
	// 		!calculator.Equal(v.currency.Exponent(), comparator.Exponent()) ||
	// 		!calculator.Equal(calculator.ComputeBase(v.currency.Base()), comparatorBase) {
	// 		return false
	// 	}
	// }

	return true
}
