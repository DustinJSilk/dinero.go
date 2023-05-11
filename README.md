# dinero.go

This is a port of the (amazing) [v2 dinero.js](https://github.com/dinerojs/dinero.js/) library written in Go.
A JSON Dinero snapshot can be unmarshalled directly into a Go Dinero, making a dinero object interchangable across a JS/Node frontend, a Go backend, a NoSql storage layer.

The API is almost identical to it's JavaScript counterpart. Please refer to the JS documentation for more information.

This library has zero dependencies and all tests from the JS implementation have been copied.

It doesn't yet support multiple base currencies.

## Usage

Dinero.go supports using any number type as its underlying value (int, int32, math/big etc).

For convenience, dinero.go provides access to quickly create `int` based Dineros as follows:

```go
package main

import (
	"github.com/DustinJSilk/dinero/dinero"
	"github.com/DustinJSilk/dinero/currency"
)

func main() {
  myDinero := dinero.NewDinero(1000, currency.USD)
}
```

Dinero.go also provides convinience methods `int` based rounding functions, for example:
`divide.DownInt` can be used to round int dineros. For custom types, you will need to pass the
type to the divider struct and ceate an instance of it: `divide.Down[int32]{}`.

You can also create your own custom type dineros with a custom calculator by implementing the
`CalculatorCore[T any]` interface.

```go
type CalculatorCore[T any] interface {
	Add(augend, addend T) T
	Compare(a, b T) CompareResult
	Decrement(value T) T
	Increment(value T) T
	IntegerDivide(dividend, divisor T) (T, error)
	Modulo(dividend, divisor T) (T, error)
	Multiply(multiplicand, multiplier T) T
	Power(base, exponent T) T
	Subtract(minuend, subtrahend T) T
	Zero() T
}
```

You can then use the new calculator by passing the calculator to the NewDineroWithOptions function:

```go
package main

import (
	"github.com/DustinJSilk/dinero/dinero"
	"github.com/DustinJSilk/dinero/currency"
)

var int32Calculator = calculator.NewCalculator(/* pass in your custom CalculatorCore */)
var USD = currency.NewCurrency[int32]("USD", 10, 2)

func NewInt32Dinero(amount int32, currency currency.Currency[int32]) dinero.Dinero[int32] {
  return dinero.NewDineroWithOptions(1000, currency, currency.Exponent, int32Calculator)
}

func main() {
  myInt32Dinero := NewInt32Dinero(1000, USD)
}
```
