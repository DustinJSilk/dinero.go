package integer

// Returns the product of two numbers.
// https://stackoverflow.com/a/71289792
func (c Calculator) Power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	y := c.Power(base, exponent/2)
	if exponent%2 == 0 {
		return y * y
	}
	return base * y * y
}
