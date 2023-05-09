package types

type CompareResult int8

var (
	LT CompareResult = -1
	EQ CompareResult = 0
	GT CompareResult = 1
)

type CalculatorCore[T any] interface {
	Add(augend, addend T) T
	Compare(a, b T) CompareResult
	Decrement(value T) T
	Increment(value T) T
	IntegerDivide(dividend, divisor T) (T, error)
	Modulo(dividend, divisor T) (T, error)
	Multiply(multiplicand, multiplier T) T
	One() T
	Power(base, exponent T) T
	Subtract(minuend, subtrahend T) T
	Zero() T
}

type Calculator[T any] interface {
	CalculatorCore[T]

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
	Sign(value T) T
}
