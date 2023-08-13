package big

import "math/big"

/*
Returns an incremented number.

@param value - The number to increment.

@returns The incremented number.
*/
func (Calculator) Increment(value *big.Int) *big.Int {
	return big.NewInt(0).Add(value, big.NewInt(1))
}
