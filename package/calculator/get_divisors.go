package calculator

func (c Calculator[T]) GetDivisors(bases ...T) []T {
	divisors := make([]T, len(bases))
	one := c.One()

	for i := range bases {
		divisor := one
		for _, d := range bases[i:] {
			divisor = c.Multiply(divisor, d)
		}

		divisors[i] = divisor
	}

	return divisors
}
