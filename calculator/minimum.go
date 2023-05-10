package calculator

// Returns the lowest value.
func (c calculator[T]) Minimum(values ...T) T {
	out := values[0]

	for _, v := range values {
		if c.LessThan(v, out) {
			out = v
		}
	}

	return out
}
