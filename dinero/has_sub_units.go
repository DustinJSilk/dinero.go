package dinero

func (d Dinero[T]) HasSubUnits() bool {
	c := d.calc()
	base := c.ComputeBase(d.Currency.Base)

	remainder, err := c.Modulo(d.Amount, c.Power(base, d.Scale))
	if err != nil {
		return false
	}

	return !c.Equal(remainder, c.Zero())
}
