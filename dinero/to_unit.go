package dinero

func (d Dinero[T]) ToUnit() ([]T, error) {
	c := d.calc()
	base := d.Currency.Base
	divisor := c.Power(base, d.Scale)

	quotient, err := c.IntegerDivide(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	remainder, err := c.Modulo(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	return []T{quotient, remainder}, nil
}
