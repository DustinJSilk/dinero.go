package currency

type Currency[T any] interface {
	Code() string
	Base() T
	Exponent() T
}

type currency[T any] struct {
	code     string
	base     T
	exponent T
}

func (c currency[T]) Code() string {
	return c.code
}

func (c currency[T]) Base() T {
	return c.base
}

func (c currency[T]) Exponent() T {
	return c.exponent
}

func NewCurrency[T any](code string, base T, exponent T) Currency[T] {
	return currency[T]{
		code:     code,
		base:     base,
		exponent: exponent,
	}
}
