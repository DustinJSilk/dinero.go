package calculator

import "dinero.go/types"

// Returns true if the subject is greater than or equal to the comparator.
func (c Calculator[T]) GreaterThanOrEqual(subject T, comparator T) bool {
	result := c.Compare(subject, comparator)
	return result == types.GT || result == types.EQ
}
