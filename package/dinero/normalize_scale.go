package dinero

import (
	"dinero.go/package/divide"
)

func NormalizeScale[T any](dineros ...Dinero[T]) ([]Dinero[T], error) {
	scales := make([]T, len(dineros))
	for i, v := range dineros {
		scales[i] = v.scale
	}

	calculator := dineros[0].calculator
	highestScale := calculator.Maximum(scales...)

	out := make([]Dinero[T], len(dineros))
	for i, v := range dineros {
		if calculator.Equal(highestScale, v.scale) {
			out[i] = v
			continue
		}
		newDinero, _ := v.TransformScale(highestScale, divide.Down[T]{})
		out[i] = newDinero
	}

	return out, nil
}

func (d Dinero[T]) NormalizeScaleWith(dineros ...Dinero[T]) ([]Dinero[T], error) {
	return NormalizeScale(append(dineros, d)...)
}
