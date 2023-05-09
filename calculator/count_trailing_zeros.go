package calculator

func (c Calculator[T]) CountTrailingZeros(value T, base T) (T, error) {
	zero := c.Zero()

	if c.Equal(value, zero) {
		return zero, nil
	}

	i := zero
	temp := value

	for {
		remainder, err := c.Modulo(temp, base)
		if err != nil {
			return zero, err
		}

		if !c.Equal(remainder, zero) {
			break
		}

		temp, err = c.IntegerDivide(temp, base)
		if err != nil {
			return zero, err
		}

		i = c.Increment(i)
	}

	return i, nil
}
