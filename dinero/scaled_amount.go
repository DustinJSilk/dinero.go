package dinero

// Used to create decimal values without floats. 1.89 is expressed as { Amount: 189, Scale: 2 }.
type ScaledAmount[T any] struct {
	Amount T
	Scale  T
}
