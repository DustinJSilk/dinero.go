package integer

import "dinero.go/types"

func (Calculator) Compare(a, b int) types.CompareResult {
	switch {
	case a < b:
		return types.LT
	case a > b:
		return types.GT
	default:
		return types.EQ
	}
}
