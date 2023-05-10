package calculator

// Returns true if the subject is less than the comparator.
func (c calculator[T]) LessThan(subject, comparator T) bool {
	return c.Compare(subject, comparator) == LT
}
