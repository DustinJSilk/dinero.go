package big

import (
	"math/big"

	"github.com/DustinJSilk/dinero.go/calculator"
)

func (Calculator) Compare(a, b *big.Int) calculator.CompareResult {
	comp := a.Cmp(b)
	switch comp {
	case -1:
		return calculator.LT
	case 1:
		return calculator.GT
	default:
		return calculator.EQ
	}
}
