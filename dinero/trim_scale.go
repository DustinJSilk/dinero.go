package dinero

// Trim a Dinero object's scale as much as possible, down to the currency exponent.
func (d Dinero[T]) TrimScale() (Dinero[T], error) {
	base := d.Calculator.ComputeBase(d.Currency.Base)
	trailingZerosLength, err := d.Calculator.CountTrailingZeros(d.Amount, base)
	if err != nil {
		return Dinero[T]{}, err
	}

	difference := d.Calculator.Subtract(d.Scale, trailingZerosLength)
	newScale := d.Calculator.Maximum(difference, d.Currency.Exponent)

	if d.Calculator.Equal(newScale, d.Scale) {
		return d, nil
	}

	return d.TransformScale(newScale, nil)
}
