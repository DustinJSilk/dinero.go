package dinero

import (
	"dinero.go/divide"
)

// Transform a Dinero object to a new scale.
//
// When transforming to a higher scale, the internal amount value increases by orders of magnitude.
// If you're using the default Dinero implementation (with the int calculator), be careful not to
// exceed the minimum and maximum safe integers.
//
// When transforming to a smaller scale, the amount loses precision. By default, the function rounds
// down the amount when passing nil as the divider. You can specify how to round by passing a custom
// divide function.
//
// For convenience, Dinero.go provides the following divide functions: up, down, halfUp, halfDown,
// halfOdd, halfEven (bankers rounding), halfTowardsZero, and halfAwayFromZero.
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
