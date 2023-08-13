package big

import "math/big"

/*
Returns the difference between two numbers.

@param minuend - The number to subtract from.
@param subtrahend - The number to subtract.

@returns The difference of the two numbers.
*/
func (Calculator) Subtract(minuend, subtrahend *big.Int) *big.Int {
	return big.NewInt(0).Sub(minuend, subtrahend)
}
