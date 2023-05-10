package dinero

import (
	"dinero.go/divide"
)

// Transform a Dinero object to a new scale.
// Passing nil to the divider will default to rounding Down.
func (d Dinero[T]) TransformScale(newScale T, divider divide.Divider[T]) (Dinero[T], error) {
	if d.calculator.Equal(d.Scale, newScale) {
		return NewDineroWithOptions(d.Amount, d.Currency, d.Scale, d.calculator), nil
	}

	isLarger := d.calculator.GreaterThan(newScale, d.Scale)
	base := d.calculator.ComputeBase(d.Currency.Base)
	var newAmount T

	if isLarger {
		factor := d.calculator.Power(base, d.calculator.Subtract(newScale, d.Scale))
		newAmount = d.calculator.Multiply(d.Amount, factor)
	} else {
		factor := d.calculator.Power(base, d.calculator.Subtract(d.Scale, newScale))

		if divider == nil {
			divider = divide.Down[T]{}
		}

		amount, err := divider.Divide(d.Amount, factor, d.calculator)
		if err != nil {
			return Dinero[T]{}, err
		}
		newAmount = amount
	}

	return NewDineroWithOptions(newAmount, d.Currency, newScale, d.calculator), nil
}
