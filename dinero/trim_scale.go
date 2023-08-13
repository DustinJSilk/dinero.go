package dinero

// Trim a Dinero object's scale as much as possible, down to the currency exponent.
func (d Dinero[T]) TrimScale() (Dinero[T], error) {
	c := d.Calculator()
	base := c.ComputeBase(d.Currency.Base)
	trailingZerosLength, err := c.CountTrailingZeros(d.Amount, base)
	if err != nil {
		return Dinero[T]{}, err
	}

	difference := c.Subtract(d.Scale, trailingZerosLength)
	newScale := c.Maximum(difference, d.Currency.Exponent)

	if c.Equal(newScale, d.Scale) {
		return d, nil
	}

	return d.TransformScale(newScale, nil)
}
