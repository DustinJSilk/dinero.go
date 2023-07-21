package dinero

func (d Dinero[T]) HasSubUnits() bool {
	base := d.Calculator.ComputeBase(d.Currency.Base)

	remainder, err := d.Calculator.Modulo(d.Amount, d.Calculator.Power(base, d.Scale))
	if err != nil {
		return false
	}

	return !d.Calculator.Equal(remainder, d.Calculator.Zero())
}
