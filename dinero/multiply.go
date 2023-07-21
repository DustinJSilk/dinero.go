package dinero

// Multiply the passed Dinero object.
// To multiply by a fraction, use MultiplyScaled.
func (d Dinero[T]) Multiply(multiplier T) Dinero[T] {
	amount := d.Calculator.Multiply(d.Amount, multiplier)
	return NewDineroWithOptions(amount, d.Currency, d.Scale, d.Calculator)
}

// Multiply the passed Dinero object by a ScaledAmount.
// To multiply by 2.1, you would pass { Amount: 21, Scale: 1 }. When using scaled amounts,
// the function converts the returned objects to the safest scale.
func (d Dinero[T]) MultiplyScaled(multiplier ScaledAmount[T]) (Dinero[T], error) {
	newScale := d.Calculator.Add(d.Scale, multiplier.Scale)

	return NewDineroWithOptions(
		d.Calculator.Multiply(d.Amount, multiplier.Amount),
		d.Currency,
		newScale,
		d.Calculator,
	).TransformScale(newScale, nil)
}
