package big_test

import (
	"math/big"
	"testing"

	bigcalc "github.com/DustinJSilk/dinero.go/calculator/big"
)

func TestZero(t *testing.T) {
	c := bigcalc.Calculator{}
	zero := c.Zero()

	if zero.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expected: 0, got: %v", zero)
	}
}
