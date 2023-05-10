package dinero

// Used to create decimal values without floats. 1.89 is expressed as { Amount: 189, Scale: 2 }.
type ScaledAmount[T any] struct {
	amount T
	scale  T
}

func NewScaledAmount[T any](amount, scale T) ScaledAmount[T] {
	return ScaledAmount[T]{
		amount: amount,
		scale:  scale,
	}
}

func (s ScaledAmount[T]) Amount() T {
	return s.amount
}

func (s ScaledAmount[T]) Scale() T {
	return s.scale
}
