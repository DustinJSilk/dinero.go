package divide

import (
	"github.com/DustinJSilk/dinero.go/calculator"
)

type Divider[T any] interface {
	Divide(amount T, factor T, calculator calculator.Calculator[T]) (T, error)
}
