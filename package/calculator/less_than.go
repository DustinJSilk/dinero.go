package calculator

import "dinero.go/package/types"

// Returns true if the subject is less than the comparator.
func (c Calculator[T]) LessThan(subject, comparator T) bool {
	return c.Compare(subject, comparator) == types.LT
}
