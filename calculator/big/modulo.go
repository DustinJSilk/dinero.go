package big

import (
	"fmt"
	"math/big"
)

// Returns the modulous of two numbers.
func (Calculator) Modulo(dividend, divisor *big.Int) (*big.Int, error) {
	if len(divisor.Bits()) == 0 {
		return nil, fmt.Errorf("modulo by zero")
	}

	_, rem := big.NewInt(0).QuoRem(dividend, divisor, big.NewInt(0))
	return rem, nil
}
