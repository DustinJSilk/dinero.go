package dinero

import (
	"fmt"

	"github.com/DustinJSilk/dinero.go/currency"
)

// Convert a Dinero object from a currency to another.
func (d Dinero[T]) Convert(currency currency.Currency[T], rates map[string]ScaledAmount[T]) (Dinero[T], error) {
	rate, ok := rates[currency.Code]
	if !ok {
		return Dinero[T]{}, fmt.Errorf("missing currency rate")
	}

	c := d.Calculator()
	newScale := c.Add(d.Scale, rate.Scale)

	return NewDineroWithOptions(
		c.Multiply(d.Amount, rate.Amount),
		currency,
		newScale,
		c,
	).TransformScale(c.Maximum(newScale, currency.Exponent), nil)
}
