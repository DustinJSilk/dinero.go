package big

import "math/big"

// Returns the product of two numbers.
func (c Calculator) Power(base, exponent *big.Int) *big.Int {
	return big.NewInt(0).Exp(base, exponent, nil)
}
