package calculator

import "dinero.go/package/types"

// Returns true if the subject is greater than the comparator.
func (c Calculator[T]) GreaterThan(subject T, comparator T) bool {
	return c.Compare(subject, comparator) == types.GT
}
