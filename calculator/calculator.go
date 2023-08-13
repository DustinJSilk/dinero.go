package calculator

type Core[T any] interface {
	Add(augend, addend T) T
	Compare(a, b T) CompareResult
	Decrement(value T) T
	Increment(value T) T
	IntegerDivide(dividend, divisor T) (T, error)
	Modulo(dividend, divisor T) (T, error)
	Multiply(multiplicand, multiplier T) T
	Power(base, exponent T) T
	Subtract(minuend, subtrahend T) T
	ToInt(v T) int
	ToString(v T) string
	Zero() T
}

type Calculator[T any] interface {
	Core[T]

	Absolute(value T) T
	ComputeBase(base ...T) T
	CountTrailingZeros(value T, base T) (T, error)
	Distribute(value T, ratios ...T) []T
	Equal(subject, comparator T) bool
	GetDivisors(bases ...T) []T
	GreaterThan(subject T, comparator T) bool
	GreaterThanOrEqual(subject T, comparator T) bool
	IsEven(value T) bool
	IsHalf(value T, total T) bool
	LessThan(subject T, comparator T) bool
	LessThanOrEqual(subject T, comparator T) bool
	Maximum(values ...T) T
	Minimum(values ...T) T
	One() T
	Sign(value T) T
	Ten() T
}

type calculator[T any] struct {
	core Core[T]
	one  T
	ten  T
}

func NewCalculator[T any](core Core[T]) calculator[T] {
	one := core.Increment(core.Zero())
	ten := core.Zero()
	for i := 0; i < 10; i++ {
		ten = core.Increment(ten)
	}

	return calculator[T]{core, one, ten}
}

func (c calculator[T]) Add(augend, addend T) T {
	return c.core.Add(augend, addend)
}

func (c calculator[T]) Compare(a, b T) CompareResult {
	return c.core.Compare(a, b)
}

func (c calculator[T]) Decrement(value T) T {
	return c.core.Decrement(value)
}

func (c calculator[T]) Increment(value T) T {
	return c.core.Increment(value)
}

func (c calculator[T]) IntegerDivide(dividend, divisor T) (T, error) {
	return c.core.IntegerDivide(dividend, divisor)
}

func (c calculator[T]) Modulo(dividend, divisor T) (T, error) {
	return c.core.Modulo(dividend, divisor)
}

func (c calculator[T]) Multiply(multiplicand, multiplier T) T {
	return c.core.Multiply(multiplicand, multiplier)
}

func (c calculator[T]) One() T {
	return c.one
}

func (c calculator[T]) Power(base, exponent T) T {
	return c.core.Power(base, exponent)
}

func (c calculator[T]) Subtract(minuend, subtrahend T) T {
	return c.core.Subtract(minuend, subtrahend)
}

func (c calculator[T]) Ten() T {
	return c.ten
}

func (c calculator[T]) ToInt(v T) int {
	return c.core.ToInt(v)
}

func (c calculator[T]) ToString(v T) string {
	return c.core.ToString(v)
}

func (c calculator[T]) Zero() T {
	return c.core.Zero()
}
