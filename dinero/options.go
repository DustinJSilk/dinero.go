package dinero

import "github.com/DustinJSilk/dinero.go/currency"

type Transformer[T any] func(value string, currency currency.Currency[T]) string

type Options[T any] struct {
	transformer Transformer[T]
}

type Option[T any] func(*Options[T])

func WithTransformer[T any](t Transformer[T]) Option[T] {
	return func(o *Options[T]) {
		o.transformer = t
	}
}
