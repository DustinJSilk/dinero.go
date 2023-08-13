package big

import "math/big"

/*
Returns the product of two numbers.

@param multiplicand - The number to multiply.
@param multiplier - The number to multiply with.

@returns The product of the two numbers.
*/
func (Calculator) Multiply(multiplicand, multiplier *big.Int) *big.Int {
	return big.NewInt(0).Mul(multiplicand, multiplier)
}
