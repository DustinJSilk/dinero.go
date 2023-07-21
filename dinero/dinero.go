package dinero

import (
	"math/big"

	"github.com/DustinJSilk/dinero.go/calculator"
	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
	"github.com/DustinJSilk/dinero.go/calculator/integer"
	"github.com/DustinJSilk/dinero.go/currency"
)

type Dinero[T any] struct {
	Amount     T                        `json:"amount"`
	Currency   currency.Currency[T]     `json:"currency"`
	Scale      T                        `json:"scale"`
	Calculator calculator.Calculator[T] `json:"-"`
}

var IntCalculator = calculator.NewCalculator[int](integer.Calculator{})
var BigIntCalculator = calculator.NewCalculator[*big.Int](bigcalc.Calculator{})

func NewDinero(amount int, currency currency.Currency[int]) Dinero[int] {
	dinero := Dinero[int]{
		Amount:     amount,
		Scale:      currency.Exponent,
		Currency:   currency,
		Calculator: IntCalculator,
	}

	return dinero
}

func NewDineroWithScale(amount int, currency currency.Currency[int], scale int) Dinero[int] {
	dinero := Dinero[int]{
		Amount:     amount,
		Scale:      scale,
		Currency:   currency,
		Calculator: calculator.NewCalculator[int](integer.Calculator{}),
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
		Calculator: calculator,
	}

	return dinero
}

func NewBigDinero(amount int64, currency currency.Currency[*big.Int]) Dinero[*big.Int] {
	return Dinero[*big.Int]{
		Amount:     big.NewInt(amount),
		Currency:   currency,
		Calculator: BigIntCalculator,
		Scale:      currency.Exponent,
	}
}

func NewBigDineroWithScale(amount int64, currency currency.Currency[*big.Int], scale int64) Dinero[*big.Int] {
	return Dinero[*big.Int]{
		Amount:     big.NewInt(amount),
		Currency:   currency,
		Scale:      big.NewInt(scale),
		Calculator: BigIntCalculator,
	}
}
