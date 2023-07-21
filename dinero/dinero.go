package dinero

import (
	"github.com/DustinJSilk/dinero.go/calculator"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
	"github.com/DustinJSilk/dinero.go/currency"
)

type Dinero[T any] struct {
	Amount     T                        `json:"amount"`
	Currency   currency.Currency[T]     `json:"currency"`
	Scale      T                        `json:"scale"`
	calculator calculator.Calculator[T] `json:"-"`
}

var IntCalculator = calculator.NewCalculator[int](integer.Calculator{})

func NewDinero(amount int, currency currency.Currency[int]) Dinero[int] {
	dinero := Dinero[int]{
		Amount:     amount,
		Scale:      currency.Exponent,
		Currency:   currency,
		calculator: IntCalculator,
	}

	return dinero
}

func NewDineroWithScale(amount int, currency currency.Currency[int], scale int) Dinero[int] {
	dinero := Dinero[int]{
		Amount:     amount,
		Scale:      scale,
		Currency:   currency,
		calculator: calculator.NewCalculator[int](integer.Calculator{}),
	}

	return dinero
}

func NewDineroWithOptions[T any](
	amount T,
	currency currency.Currency[T],
	scale T,
	calculator calculator.Calculator[T],
) Dinero[T] {
	dinero := Dinero[T]{
		Amount:     amount,
		Currency:   currency,
		Scale:      scale,
		calculator: calculator,
	}

	return dinero
}
