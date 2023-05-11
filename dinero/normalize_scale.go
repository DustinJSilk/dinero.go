package dinero

import (
	"dinero.go/divide"
)

func NormalizeScale[T any](dineros ...Dinero[T]) []Dinero[T] {
	scales := make([]T, len(dineros))
	for i, v := range dineros {
		scales[i] = v.Scale
	}

	calculator := dineros[0].calculator
	highestScale := calculator.Maximum(scales...)

	out := make([]Dinero[T], len(dineros))
	for i, v := range dineros {
		if calculator.Equal(highestScale, v.Scale) {
			out[i] = v
			continue
		}
		newDinero, _ := v.TransformScale(highestScale, divide.Down[T]{})
		out[i] = newDinero
	}

	return out
}
