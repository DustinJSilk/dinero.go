package dinero

import (
	"dinero.go/divide"
)

// Normalize a set of Dinero objects to the highest scale of the set.
//
// Normalizing to a higher scale means that the internal amount value increases by orders of
// magnitude. If you're using the default Dinero implementation (with the int calculator), be
// careful not to exceed the minimum and maximum safe integers.
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
