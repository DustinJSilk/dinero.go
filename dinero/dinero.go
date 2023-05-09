package dinero

import (
	"dinero.go/calculator"
	"dinero.go/calculator/integer"
	"dinero.go/currency"
	"dinero.go/types"
)

type Dinero[T any] struct {
	amount     T
	currency   currency.Currency[T]
	scale      T
	calculator types.Calculator[T]
}

func NewDinero(amount int, currency currency.Currency[int]) Dinero[int] {
	dinero := Dinero[int]{
		amount:     amount,
		scale:      currency.Exponent(),
		currency:   currency,
		calculator: calculator.NewCalculator[int](integer.Calculator{}),
	}

	return dinero
}

func NewDineroWithScale(amount int, currency currency.Currency[int], scale int) Dinero[int] {
	dinero := Dinero[int]{
		amount:     amount,
		scale:      scale,
		currency:   currency,
		calculator: calculator.NewCalculator[int](integer.Calculator{}),
	}

	return dinero
}

func NewDineroWithOptions[T any](
	amount T,
	currency currency.Currency[T],
	scale T,
	calculator types.Calculator[T],
) Dinero[T] {
	dinero := Dinero[T]{
		amount:     amount,
		currency:   currency,
		scale:      scale,
		calculator: calculator,
	}

	return dinero
}
