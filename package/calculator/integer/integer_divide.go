package integer

import "fmt"

// Returns the quotient of two numbers with no fractional part.
func (Calculator) IntegerDivide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("divide by zero")
	}

	return dividend / divisor, nil
}
