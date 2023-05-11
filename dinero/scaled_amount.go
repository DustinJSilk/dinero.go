package dinero

// Used to create decimal values without floats. 1.89 is expressed as { Amount: 189, Scale: 2 }.
type ScaledAmount[T any] struct {
	Amount T `json:"amount"`
	Scale  T `json:"scale"`
}

func NewScaledAmount[T any](amount, scale T) ScaledAmount[T] {
	return ScaledAmount[T]{
		Amount: amount,
		Scale:  scale,
	}
}
