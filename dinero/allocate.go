package dinero

import (
	"fmt"
)

func unsafeAllocate[T any](dinero Dinero[T], ratios []T) []Dinero[T] {
	shares := dinero.calculator.Distribute(dinero.amount, ratios...)

	dineros := make([]Dinero[T], len(shares))
	for i, v := range shares {
		dineros[i] = NewDineroWithOptions(v, dinero.currency, dinero.scale, dinero.calculator)
	}

	return dineros
}

// Distribute the amount of a Dinero object across a list of ratios.
// Unlike in dinero.js, this function does not support distributing across differently scaled ratios.
func (d Dinero[T]) Allocate(ratios ...T) ([]Dinero[T], error) {
	if len(ratios) == 0 {
		return nil, fmt.Errorf("missing ratios")
	}

	zero := d.calculator.Zero()
	hasOnlyPositiveRatios := true
	hasOneNonZeroRatio := false
	for _, v := range ratios {
		if d.calculator.LessThan(v, zero) {
			hasOnlyPositiveRatios = false
			break
		}
		if d.calculator.GreaterThan(v, zero) {
			hasOneNonZeroRatio = true
		}
	}

	if !hasOnlyPositiveRatios {
		return nil, fmt.Errorf("ratios must be positive")
	}

	if !hasOneNonZeroRatio {
		return nil, fmt.Errorf("must have at least 1 non-zero ratio")
	}

	return unsafeAllocate(d, ratios), nil
}
