package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/currency"
	"dinero.go/dinero"
	"dinero.go/divide"
)

func TestTransformScale(t *testing.T) {
	type test struct {
		description string
		input       dinero.Dinero[int]
		scale       int
		expect      dinero.Dinero[int]
		divide      divide.Divider[int]
	}

	abc := currency.NewCurrency("ABC", 6, 1)

	tests := []test{
		{
			description: "returns a new Dinero object with a new scale and a converted amount",
			input:       dinero.NewDineroWithScale(500, currency.USD, 2),
			scale:       4,
			expect:      dinero.NewDineroWithScale(50000, currency.USD, 4),
		},
		{
			description: "returns a new Dinero object with a new scale and a converted, rounded down amount",
			input:       dinero.NewDineroWithScale(14270, currency.USD, 2),
			scale:       0,
			expect:      dinero.NewDineroWithScale(142, currency.USD, 0),
		},
		{
			description: "converts between scales correctly",
			input:       dinero.NewDineroWithScale(333336, currency.USD, 5),
			scale:       2,
			expect:      dinero.NewDineroWithScale(333, currency.USD, 2),
		},
		{
			description: "converts from long initial scales correctly",
			input:       dinero.NewDineroWithScale(3333333336, currency.USD, 9),
			scale:       2,
			expect:      dinero.NewDineroWithScale(333, currency.USD, 2),
		},
		{
			description: "ignores equal scales",
			input:       dinero.NewDineroWithScale(5000, currency.USD, 3),
			scale:       3,
			expect:      dinero.NewDineroWithScale(5000, currency.USD, 3),
		},
		{
			description: "uses the provided `up` divide function",
			input:       dinero.NewDineroWithScale(10455, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1046, currency.USD, 2),
			divide:      divide.UpInt,
		},
		{
			description: "uses the provided `down` divide function",
			input:       dinero.NewDineroWithScale(10455, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1045, currency.USD, 2),
			divide:      divide.DownInt,
		},
		{
			description: "uses the provided `halfOdd` divide function A",
			input:       dinero.NewDineroWithScale(10415, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1041, currency.USD, 2),
			divide:      divide.HalfOddInt,
		},
		{
			description: "uses the provided `halfOdd` divide function B",
			input:       dinero.NewDineroWithScale(10425, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1043, currency.USD, 2),
			divide:      divide.HalfOddInt,
		},
		{
			description: "uses the provided `halfEven` divide function A",
			input:       dinero.NewDineroWithScale(10425, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1042, currency.USD, 2),
			divide:      divide.HalfEvenInt,
		},
		{
			description: "uses the provided `halfEven` divide function B",
			input:       dinero.NewDineroWithScale(10435, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1044, currency.USD, 2),
			divide:      divide.HalfEvenInt,
		},
		{
			description: "uses the provided `halfDown` divide function A",
			input:       dinero.NewDineroWithScale(10455, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1045, currency.USD, 2),
			divide:      divide.HalfDownInt,
		},
		{
			description: "uses the provided `halfDown` divide function B",
			input:       dinero.NewDineroWithScale(10456, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1046, currency.USD, 2),
			divide:      divide.HalfDownInt,
		},
		{
			description: "uses the provided `halfUp` divide function A",
			input:       dinero.NewDineroWithScale(10454, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1045, currency.USD, 2),
			divide:      divide.HalfUpInt,
		},
		{
			description: "uses the provided `halfUp` divide function B",
			input:       dinero.NewDineroWithScale(10455, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1046, currency.USD, 2),
			divide:      divide.HalfUpInt,
		},
		{
			description: "uses the provided `HalfTowardsZero` divide function",
			input:       dinero.NewDineroWithScale(10415, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1041, currency.USD, 2),
			divide:      divide.HalfTowardsZeroInt,
		},
		{
			description: "uses the provided `HalfAwayFromZero` divide function",
			input:       dinero.NewDineroWithScale(10415, currency.USD, 3),
			scale:       2,
			expect:      dinero.NewDineroWithScale(1042, currency.USD, 2),
			divide:      divide.HalfAwayFromZeroInt,
		},
		// non-decimal currencies
		{
			description: "returns a new Dinero object with a new scale and a converted amount",
			input:       dinero.NewDinero(5, currency.MGA),
			scale:       2,
			expect:      dinero.NewDineroWithScale(25, currency.MGA, 2),
		},
		{
			description: "returns a new Dinero object with a new scale and a converted, rounded down amount",
			input:       dinero.NewDineroWithScale(26, currency.MGA, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(5, currency.MGA, 1),
		},
		{
			description: "uses the provided `up` divide function",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.UpInt,
		},
		{
			description: "uses the provided `down` divide function",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(5, abc, 1),
			divide:      divide.DownInt,
		},
		{
			description: "uses the provided `halfOdd` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(5, abc, 1),
			divide:      divide.HalfOddInt,
		},
		{
			description: "uses the provided `halfOdd` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(7, abc, 1),
			divide:      divide.HalfOddInt,
		},
		{
			description: "uses the provided `halfEven` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfEvenInt,
		},
		{
			description: "uses the provided `halfEven` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfEvenInt,
		},
		{
			description: "uses the provided `halfDown` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(5, abc, 1),
			divide:      divide.HalfDownInt,
		},
		{
			description: "uses the provided `halfDown` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfDownInt,
		},
		{
			description: "uses the provided `halfUp` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfUpInt,
		},
		{
			description: "uses the provided `halfUp` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(7, abc, 1),
			divide:      divide.HalfUpInt,
		},
		{
			description: "uses the provided `halfTowardsZero` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(5, abc, 1),
			divide:      divide.HalfTowardsZeroInt,
		},
		{
			description: "uses the provided `halfTowardsZero` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfTowardsZeroInt,
		},
		{
			description: "uses the provided `halfAwayFromZero` divide function A",
			input:       dinero.NewDineroWithScale(33, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(6, abc, 1),
			divide:      divide.HalfAwayFromZeroInt,
		},
		{
			description: "uses the provided `halfAwayFromZero` divide function B",
			input:       dinero.NewDineroWithScale(39, abc, 2),
			scale:       1,
			expect:      dinero.NewDineroWithScale(7, abc, 1),
			divide:      divide.HalfAwayFromZeroInt,
		},
	}

	for _, tc := range tests {
		var got dinero.Dinero[int]
		var err error

		if tc.divide != nil {
			got, err = tc.input.TransformScale(tc.scale, tc.divide)
		} else {
			got, err = tc.input.TransformScale(tc.scale, divide.DownInt)
		}

		if err != nil {
			t.Fatalf("error: %v, %e", tc.expect, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Fatalf("expected: %v, got: %v", tc.expect, got)
		}
	}
}

func BenchmarkTransformScale(b *testing.B) {
	a := dinero.NewDineroWithScale(1000, currency.USD, 3)

	for i := 0; i < b.N; i++ {
		_, err := a.TransformScale(2, divide.DownInt)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
