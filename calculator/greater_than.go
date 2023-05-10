package calculator

// Returns true if the subject is greater than the comparator.
func (c calculator[T]) GreaterThan(subject T, comparator T) bool {
	return c.Compare(subject, comparator) == GT
}
