package calculator

// Returns the highest value.
func (c calculator[T]) Maximum(values ...T) T {
	out := values[0]

	for _, v := range values {
		if c.GreaterThan(v, out) {
			out = v
		}
	}

	return out
}
