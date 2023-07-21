package dinero

func (d Dinero[T]) ToUnit() ([]T, error) {
	base := d.Currency.Base
	divisor := d.Calculator.Power(base, d.Scale)

	quotient, err := d.Calculator.IntegerDivide(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	remainder, err := d.Calculator.Modulo(d.Amount, divisor)
	if err != nil {
		return nil, err
	}

	return []T{quotient, remainder}, nil
}
