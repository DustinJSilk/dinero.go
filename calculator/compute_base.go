package calculator

func (c calculator[T]) ComputeBase(base ...T) T {
	if len(base) == 1 {
		return base[0]
	}

	out := c.One()

	for _, b := range base {
		out = c.Multiply(out, b)
	}

	return out
}
