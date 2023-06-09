package integer_test

import (
	"testing"

	"github.com/DustinJSilk/dinero.go/calculator/integer"
)

func TestZero(t *testing.T) {
	c := integer.Calculator{}
	zero := c.Zero()

	if zero != 0 {
		t.Fatalf("expected: 0, got: %v", zero)
	}
}
