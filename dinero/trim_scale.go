package dinero

// Trim a Dinero object's scale as much as possible, down to the currency exponent.
func (d Dinero[T]) TrimScale() (Dinero[T], error) {
	base := d.calculator.ComputeBase(d.Currency.Base)
	trailingZerosLength, err := d.calculator.CountTrailingZeros(d.Amount, base)
	if err != nil {
		return Dinero[T]{}, err
	}

	difference := d.calculator.Subtract(d.Scale, trailingZerosLength)
	newScale := d.calculator.Maximum(difference, d.Currency.Exponent)

	if d.calculator.Equal(newScale, d.Scale) {
		return d, nil
	}

	return d.TransformScale(newScale, nil)
}
