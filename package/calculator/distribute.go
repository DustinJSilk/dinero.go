package calculator

func (c Calculator[T]) Distribute(value T, ratios ...T) []T {
	zero := c.Zero()
	one := c.One()

	total := zero
	for _, v := range ratios {
		total = c.Add(total, v)
	}

	if c.Equal(total, zero) {
		return ratios
	}

	remainder := value
	shares := make([]T, len(ratios))
	for i, ratio := range ratios {
		share, _ := c.IntegerDivide(c.Multiply(value, ratio), total)
		remainder = c.Subtract(remainder, share)
		shares[i] = share
	}

	isPositive := c.GreaterThanOrEqual(value, zero)
	var amount T

	if isPositive {
		amount = one
	} else {
		amount = c.Decrement(zero)
	}

	for i, v := range ratios {
		if isPositive {
			if c.LessThanOrEqual(remainder, zero) {
				break
			}
		} else {
			if c.GreaterThanOrEqual(remainder, zero) {
				break
			}
		}

		if !c.Equal(v, zero) {
			shares[i] = c.Add(shares[i], amount)
			remainder = c.Subtract(remainder, amount)
		}
	}

	return shares
}
