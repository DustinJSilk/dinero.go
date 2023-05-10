package currency

type Currency[T any] struct {
	Code     string `json:"code"`
	Base     T      `json:"base"`
	Exponent T      `json:"exponent"`
}

func NewCurrency[T any](code string, base T, exponent T) Currency[T] {
	return Currency[T]{
		Code:     code,
		Base:     base,
		Exponent: exponent,
	}
}
