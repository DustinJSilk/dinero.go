package dinero

import (
	"log"
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
	calculator calculator.Calculator[T] `json:"-"`
}

var IntCalculator = calculator.NewCalculator[int](integer.Calculator{})
var BigIntCalculator = calculator.NewCalculator[*big.Int](bigcalc.Calculator{})

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

func NewBigDinero(amount int64, currency currency.Currency[*big.Int]) Dinero[*big.Int] {
	return Dinero[*big.Int]{
		Amount:     big.NewInt(amount),
		Currency:   currency,
		calculator: BigIntCalculator,
		Scale:      currency.Exponent,
	}
}

func NewBigDineroWithScale(amount int64, currency currency.Currency[*big.Int], scale int64) Dinero[*big.Int] {
	return Dinero[*big.Int]{
		Amount:     big.NewInt(amount),
		Currency:   currency,
		Scale:      big.NewInt(scale),
		calculator: BigIntCalculator,
	}
}

func (d *Dinero[T]) WithCalculator(calculator calculator.Calculator[T]) Dinero[T] {
	return NewDineroWithOptions(d.Amount, d.Currency, d.Scale, calculator)
}

// Get the calculator or find the correct type if nil.
func (d *Dinero[T]) Calculator() calculator.Calculator[T] {
	if d.calculator != nil {
		return d.calculator
	}

	switch any(d.Amount).(type) {
	case int:
		return castIntCalculator[T](&IntCalculator)
	case *big.Int:
		return castBigCalculator[T](&BigIntCalculator)
	default:
		log.Fatal("dinero calculator not found")
		return castIntCalculator[T](&IntCalculator)
	}
}

func castIntCalculator[T any](c calculator.Calculator[int]) calculator.Calculator[T] {
	return c.(calculator.Calculator[T])
}

func castBigCalculator[T any](c calculator.Calculator[*big.Int]) calculator.Calculator[T] {
	return c.(calculator.Calculator[T])
}
