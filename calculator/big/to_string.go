package big

import "math/big"

func (Calculator) ToString(v *big.Int) string {
	return v.String()
}
