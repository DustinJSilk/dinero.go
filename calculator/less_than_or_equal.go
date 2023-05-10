package calculator

// Returns true if the subject is greater than or equal to the comparator.
func (c calculator[T]) LessThanOrEqual(subject T, comparator T) bool {
	result := c.Compare(subject, comparator)
	return result == LT || result == EQ
}
