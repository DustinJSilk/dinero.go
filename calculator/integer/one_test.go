package integer_test

import (
	"testing"

	"dinero.go/calculator/integer"
)

func TestOne(t *testing.T) {
	c := integer.Calculator{}
	one := c.One()

	if one != 1 {
		t.Fatalf("expected: 1, got: %v", one)
	}
}
