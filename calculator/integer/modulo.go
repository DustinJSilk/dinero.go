package integer

import "fmt"

// Returns the modulous of two numbers.
func (Calculator) Modulo(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("modulo by zero")
	}

	return dividend % divisor, nil
}
