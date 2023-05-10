package dinero

func (d Dinero[T]) HasSubUnits() bool {
	base := d.calculator.ComputeBase(d.Currency.Base)

	remainder, err := d.calculator.Modulo(d.Amount, d.calculator.Power(base, d.Scale))
	if err != nil {
		return false
	}

	return !d.calculator.Equal(remainder, d.calculator.Zero())
}
