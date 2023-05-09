package integer

// Returns the product of two numbers.
func (Calculator) Power(base, exponent int) int {
	return PowInt(base, exponent)
}

// https://stackoverflow.com/a/71289792
func PowInt(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInt(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}
