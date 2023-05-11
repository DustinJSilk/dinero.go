package dinero

import (
	"fmt"

	"dinero.go/currency"
)

// Convert a Dinero object from a currency to another.
func (d Dinero[T]) Convert(currency currency.Currency[T], rates map[string]ScaledAmount[T]) (Dinero[T], error) {
	rate, ok := rates[currency.Code]
	if !ok {
		return Dinero[T]{}, fmt.Errorf("missing currency rate")
	}

	newScale := d.calculator.Add(d.Scale, rate.Scale)

	return NewDineroWithOptions(
		d.calculator.Multiply(d.Amount, rate.Amount),
		currency,
		newScale,
		d.calculator,
	).TransformScale(
		d.calculator.Maximum(newScale, currency.Exponent),
		nil,
	)
}
