package calculator

import "dinero.go/types"

// Returns true if the subject is greater than or equal to the comparator.
func (c Calculator[T]) LessThanOrEqual(subject T, comparator T) bool {
	result := c.Compare(subject, comparator)
	return result == types.LT || result == types.EQ
}
