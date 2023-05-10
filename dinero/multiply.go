package dinero

// Multiply the passed Dinero object.
// To multiply by a fraction, use MultiplyScaled
func (d Dinero[T]) Multiply(multiplier T) (Dinero[T], error) {
	amount := d.calculator.Multiply(d.amount, multiplier)
	return NewDineroWithOptions(amount, d.currency, d.scale, d.calculator), nil
}

// Multiply the passed Dinero object by a ScaledAmount.
func (d Dinero[T]) MultiplyScaled(multiplier ScaledAmount[T]) (Dinero[T], error) {
	newScale := d.calculator.Add(d.scale, multiplier.Scale())

	return NewDineroWithOptions(
		d.calculator.Multiply(d.amount, multiplier.Amount()),
		d.currency,
		newScale,
		d.calculator,
	).TransformScale(newScale, nil)
}
