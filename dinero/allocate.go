package dinero

import (
	"fmt"

	"dinero.go/divide"
)

func unsafeAllocate[T any](dinero Dinero[T], ratios []T) []Dinero[T] {
	shares := dinero.calculator.Distribute(dinero.Amount, ratios...)

	dineros := make([]Dinero[T], len(shares))
	for i, v := range shares {
		dineros[i] = NewDineroWithOptions(v, dinero.Currency, dinero.Scale, dinero.calculator)
	}

	return dineros
}

// Distribute the amount of a Dinero object across a list of ratios.
// To distribute with a ratio less than 1, use the AllocateScaled function.
//
// Monetary values have indivisible units, meaning you can't always exactly split them. With
// allocate, you can split a monetary amount then distribute the remainder as evenly as possible.
// You can use percentage or ratio style for ratios: [25, 75] and [1, 3] do the same thing. You can
// also pass zero ratios (such as [0, 50, 50]). If there's a remainder to distribute, zero ratios
// are skipped and return a Dinero object with amount zero.
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

// Distribute the amount of a Dinero object across a list of scaled ratios.
func (d Dinero[T]) AllocateScaled(ratios ...ScaledAmount[T]) ([]Dinero[T], error) {
	c := d.calculator

	scales := make([]T, len(ratios))
	for i, v := range ratios {
		scales[i] = v.Scale
	}

	highestScale := c.Maximum(scales...)

	normalizedRatios := make([]T, len(ratios))
	for i, v := range ratios {
		factor := c.Zero()

		if !c.Equal(v.Scale, highestScale) {
			factor = c.Subtract(highestScale, v.Scale)
		}

		normalizedRatios[i] = c.Multiply(v.Amount, c.Power(c.Ten(), factor))
	}

	newScale := c.Add(d.Scale, highestScale)
	transformed, err := d.TransformScale(newScale, divide.Down[T]{})
	if err != nil {
		return nil, err
	}

	return transformed.Allocate(normalizedRatios...)
}
