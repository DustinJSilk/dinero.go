package integer

import (
	"dinero.go/calculator"
)

func (Calculator) Compare(a, b int) calculator.CompareResult {
	switch {
	case a < b:
		return calculator.LT
	case a > b:
		return calculator.GT
	default:
		return calculator.EQ
	}
}
