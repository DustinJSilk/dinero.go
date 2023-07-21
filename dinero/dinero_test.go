package dinero_test

import (
	"math/big"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

type BigDinero dinero.Dinero[*big.Int]

var BigUSD = currency.Currency[*big.Int]{
	Code:     "USD",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(2),
}

func NewBigDinero(amount int64, currency currency.Currency[*big.Int]) BigDinero {
	return BigDinero{
		Amount:   big.NewInt(amount),
		Currency: currency,
	}
}

func NewBigDineroWithScale(amount int64, currency currency.Currency[*big.Int], scale int64) BigDinero {
	return BigDinero{
		Amount:   big.NewInt(amount),
		Currency: currency,
		Scale:    big.NewInt(scale),
	}
}
