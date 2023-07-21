package big

import "math/big"

// Returns a value equal to 0.
func (Calculator) Zero() *big.Int {
	return big.NewInt(0)
}
