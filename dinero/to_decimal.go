package dinero

import (
	"errors"
	"fmt"
)

var ErrNonDecimalCurrency = errors.New("non-decimal currency")

func (d Dinero[T]) ToDecimal(options ...Option[T]) (string, error) {
	c := d.Calculator()

	opts := Options[T]{}
	for _, op := range options {
		op(&opts)
	}

	base := d.Currency.Base
	isBaseTenRemained, err := c.Modulo(base, c.Ten())
	if err != nil {
		return "", err
	}
	isBaseTen := c.Equal(isBaseTenRemained, c.Zero())

	if !isBaseTen {
		return "", ErrNonDecimalCurrency
	}

	units, err := d.ToUnit()
	if err != nil {
		return "", err
	}

	whole := c.ToString(units[0])
	fractional := c.ToString(c.Absolute(units[1]))

	scaleNumber := c.ToInt(d.Scale)
	decimal := fmt.Sprintf("%s.%s", whole, fmt.Sprintf("%0*s", scaleNumber, fractional))

	leadsWithZero := c.Equal(units[0], c.Zero())
	isNegative := c.LessThan(units[1], c.Zero())

	if leadsWithZero && isNegative {
		decimal = fmt.Sprintf("-%s", decimal)
	}

	if opts.transformer == nil {
		return decimal, nil
	}

	return opts.transformer(decimal, d.Currency), nil
}
