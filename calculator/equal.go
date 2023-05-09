package calculator

import (
	"dinero.go/types"
)

// Returns true if the comparator is equal to the subject.
func (c Calculator[T]) Equal(subject, comparator T) bool {
	return c.Compare(subject, comparator) == types.EQ
}
