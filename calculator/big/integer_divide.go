package big

import (
	"fmt"
	"math/big"
)

// Returns the quotient of two numbers with no fractional part.
func (Calculator) IntegerDivide(dividend, divisor *big.Int) (*big.Int, error) {
	if len(divisor.Bits()) == 0 {
		return nil, fmt.Errorf("divide by zero")
	}

	return big.NewInt(0).Quo(dividend, divisor), nil
}
