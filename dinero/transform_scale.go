package dinero

import (
	"dinero.go/divide"
)

// Transform a Dinero object to a new scale.
// Passing nil to the divider will default to rounding Down
func (d Dinero[T]) TransformScale(newScale T, divider divide.Divider[T]) (Dinero[T], error) {
	if d.calculator.Equal(d.scale, newScale) {
		return NewDineroWithOptions(d.amount, d.currency, d.scale, d.calculator), nil
	}

	isLarger := d.calculator.GreaterThan(newScale, d.scale)
	base := d.calculator.ComputeBase(d.currency.Base())
	var newAmount T

	if isLarger {
		factor := d.calculator.Power(base, d.calculator.Subtract(newScale, d.scale))
		newAmount = d.calculator.Multiply(d.amount, factor)
	} else {
		factor := d.calculator.Power(base, d.calculator.Subtract(d.scale, newScale))

		if divider == nil {
			divider = divide.Down[T]{}
		}

		amount, err := divider.Divide(d.amount, factor, d.calculator)
		if err != nil {
			return Dinero[T]{}, err
		}
		newAmount = amount
	}

	return NewDineroWithOptions(newAmount, d.currency, newScale, d.calculator), nil
}
