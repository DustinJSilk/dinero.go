package dinero_test

import (
	"math/big"

	"github.com/DustinJSilk/dinero.go/currency"
)

var BigUSD = currency.Currency[*big.Int]{
	Code:     "USD",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(2),
}

var BigEUR = currency.Currency[*big.Int]{
	Code:     "EUR",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(2),
}

var BigMGA = currency.Currency[*big.Int]{
	Code:     "MGA",
	Base:     big.NewInt(5),
	Exponent: big.NewInt(1),
}

var BigMRU = currency.Currency[*big.Int]{
	Code:     "MRU",
	Base:     big.NewInt(5),
	Exponent: big.NewInt(1),
}

var BigIQD = currency.Currency[*big.Int]{
	Code:     "IQD",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(3),
}
