package big

import "math/big"

func (Calculator) ToInt(v *big.Int) int {
	return int(v.Int64())
}
