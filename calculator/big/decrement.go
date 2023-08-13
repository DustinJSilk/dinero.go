package big

import "math/big"

/*
Returns an decremented number.

@param value - The number to decrement.

@returns The decremented number.
*/
func (Calculator) Decrement(value *big.Int) *big.Int {
	return big.NewInt(0).Sub(value, big.NewInt(1))
}
