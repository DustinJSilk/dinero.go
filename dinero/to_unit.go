package dinero

func (d Dinero[T]) ToUnit() ([]T, error) {
	base := d.Currency.Base
	divisor := d.calculator.Power(base, d.Scale)

	quotient, err := d.calculator.IntegerDivide(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	remainder, err := d.calculator.Modulo(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	return []T{quotient, remainder}, nil
}
