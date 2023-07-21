package big

import "math/big"

func (Calculator) Add(augend, addend *big.Int) *big.Int {
	return big.NewInt(0).Add(augend, addend)
}
