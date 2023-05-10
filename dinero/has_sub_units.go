package dinero

func (d Dinero[T]) HasSubUnits() bool {
	base := d.calculator.ComputeBase(d.currency.Base())

	remainder, err := d.calculator.Modulo(d.amount, d.calculator.Power(base, d.scale))
	if err != nil {
		return false
	}

	return !d.calculator.Equal(remainder, d.calculator.Zero())
}
