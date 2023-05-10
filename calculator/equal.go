package calculator

// Returns true if the comparator is equal to the subject.
func (c calculator[T]) Equal(subject, comparator T) bool {
	return c.Compare(subject, comparator) == EQ
}
