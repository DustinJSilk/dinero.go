package dinero

import (
	"github.com/DustinJSilk/dinero.go/divide"
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
	c := d.Calculator()
	if c.Equal(d.Scale, newScale) {
		return NewDineroWithOptions(d.Amount, d.Currency, d.Scale, c), nil
	}

	isLarger := c.GreaterThan(newScale, d.Scale)
	base := c.ComputeBase(d.Currency.Base)
	var newAmount T

	if isLarger {
		factor := c.Power(base, c.Subtract(newScale, d.Scale))
		newAmount = c.Multiply(d.Amount, factor)
	} else {
		factor := c.Power(base, c.Subtract(d.Scale, newScale))

		if divider == nil {
			divider = divide.Down[T]{}
		}

		amount, err := divider.Divide(d.Amount, factor, c)
		if err != nil {
			return Dinero[T]{}, err
		}
		newAmount = amount
	}

	return NewDineroWithOptions(newAmount, d.Currency, newScale, c), nil
}
