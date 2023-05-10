package dinero

// Multiply the passed Dinero object.
// To multiply by a fraction, use MultiplyScaled.
func (d Dinero[T]) Multiply(multiplier T) (Dinero[T], error) {
	amount := d.calculator.Multiply(d.Amount, multiplier)
	return NewDineroWithOptions(amount, d.Currency, d.Scale, d.calculator), nil
}

// Multiply the passed Dinero object by a ScaledAmount.
func (d Dinero[T]) MultiplyScaled(multiplier ScaledAmount[T]) (Dinero[T], error) {
	newScale := d.calculator.Add(d.Scale, multiplier.Scale())

	return NewDineroWithOptions(
		d.calculator.Multiply(d.Amount, multiplier.Amount()),
		d.Currency,
		newScale,
		d.calculator,
	).TransformScale(newScale, nil)
}
