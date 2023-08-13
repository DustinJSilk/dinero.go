package dinero_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
	"github.com/DustinJSilk/dinero.go/divide"
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

func TestTransformScaleBigInt(t *testing.T) {
	type test struct {
		description string
		input       dinero.Dinero[*big.Int]
		scale       *big.Int
		expect      dinero.Dinero[*big.Int]
		divide      divide.Divider[*big.Int]
	}

	abc := currency.NewCurrency("ABC", big.NewInt(6), big.NewInt(1))

	tests := []test{
		{
			description: "returns a new Dinero object with a new scale and a converted amount",
			input:       dinero.NewBigDineroWithScale(500, BigUSD, 2),
			scale:       big.NewInt(4),
			expect:      dinero.NewBigDineroWithScale(50000, BigUSD, 4),
		},
		{
			description: "returns a new Dinero object with a new scale and a converted, rounded down amount",
			input:       dinero.NewBigDineroWithScale(14270, BigUSD, 2),
			scale:       big.NewInt(0),
			expect:      dinero.NewBigDineroWithScale(142, BigUSD, 0),
		},
		{
			description: "converts between scales correctly",
			input:       dinero.NewBigDineroWithScale(333336, BigUSD, 5),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(333, BigUSD, 2),
		},
		{
			description: "converts from long initial scales correctly",
			input:       dinero.NewBigDineroWithScale(3333333336, BigUSD, 9),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(333, BigUSD, 2),
		},
		{
			description: "ignores equal scales",
			input:       dinero.NewBigDineroWithScale(5000, BigUSD, 3),
			scale:       big.NewInt(3),
			expect:      dinero.NewBigDineroWithScale(5000, BigUSD, 3),
		},
		{
			description: "uses the provided `up` divide function",
			input:       dinero.NewBigDineroWithScale(10455, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1046, BigUSD, 2),
			divide:      divide.UpBigInt,
		},
		{
			description: "uses the provided `down` divide function",
			input:       dinero.NewBigDineroWithScale(10455, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1045, BigUSD, 2),
			divide:      divide.DownBigInt,
		},
		{
			description: "uses the provided `halfOdd` divide function A",
			input:       dinero.NewBigDineroWithScale(10415, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1041, BigUSD, 2),
			divide:      divide.HalfOddBigInt,
		},
		{
			description: "uses the provided `halfOdd` divide function B",
			input:       dinero.NewBigDineroWithScale(10425, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1043, BigUSD, 2),
			divide:      divide.HalfOddBigInt,
		},
		{
			description: "uses the provided `halfEven` divide function A",
			input:       dinero.NewBigDineroWithScale(10425, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1042, BigUSD, 2),
			divide:      divide.HalfEvenBigInt,
		},
		{
			description: "uses the provided `halfEven` divide function B",
			input:       dinero.NewBigDineroWithScale(10435, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1044, BigUSD, 2),
			divide:      divide.HalfEvenBigInt,
		},
		{
			description: "uses the provided `halfDown` divide function A",
			input:       dinero.NewBigDineroWithScale(10455, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1045, BigUSD, 2),
			divide:      divide.HalfDownBigInt,
		},
		{
			description: "uses the provided `halfDown` divide function B",
			input:       dinero.NewBigDineroWithScale(10456, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1046, BigUSD, 2),
			divide:      divide.HalfDownBigInt,
		},
		{
			description: "uses the provided `halfUp` divide function A",
			input:       dinero.NewBigDineroWithScale(10454, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1045, BigUSD, 2),
			divide:      divide.HalfUpBigInt,
		},
		{
			description: "uses the provided `halfUp` divide function B",
			input:       dinero.NewBigDineroWithScale(10455, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1046, BigUSD, 2),
			divide:      divide.HalfUpBigInt,
		},
		{
			description: "uses the provided `HalfTowardsZero` divide function",
			input:       dinero.NewBigDineroWithScale(10415, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1041, BigUSD, 2),
			divide:      divide.HalfTowardsZeroBigInt,
		},
		{
			description: "uses the provided `HalfAwayFromZero` divide function",
			input:       dinero.NewBigDineroWithScale(10415, BigUSD, 3),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(1042, BigUSD, 2),
			divide:      divide.HalfAwayFromZeroBigInt,
		},
		// non-decimal currencies
		{
			description: "returns a new Dinero object with a new scale and a converted amount",
			input:       dinero.NewBigDinero(5, BigMGA),
			scale:       big.NewInt(2),
			expect:      dinero.NewBigDineroWithScale(25, BigMGA, 2),
		},
		{
			description: "returns a new Dinero object with a new scale and a converted, rounded down amount",
			input:       dinero.NewBigDineroWithScale(26, BigMGA, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(5, BigMGA, 1),
		},
		{
			description: "uses the provided `up` divide function",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.UpBigInt,
		},
		{
			description: "uses the provided `down` divide function",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(5, abc, 1),
			divide:      divide.DownBigInt,
		},
		{
			description: "uses the provided `halfOdd` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(5, abc, 1),
			divide:      divide.HalfOddBigInt,
		},
		{
			description: "uses the provided `halfOdd` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(7, abc, 1),
			divide:      divide.HalfOddBigInt,
		},
		{
			description: "uses the provided `halfEven` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfEvenBigInt,
		},
		{
			description: "uses the provided `halfEven` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfEvenBigInt,
		},
		{
			description: "uses the provided `halfDown` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(5, abc, 1),
			divide:      divide.HalfDownBigInt,
		},
		{
			description: "uses the provided `halfDown` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfDownBigInt,
		},
		{
			description: "uses the provided `halfUp` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfUpBigInt,
		},
		{
			description: "uses the provided `halfUp` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(7, abc, 1),
			divide:      divide.HalfUpBigInt,
		},
		{
			description: "uses the provided `halfTowardsZero` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(5, abc, 1),
			divide:      divide.HalfTowardsZeroBigInt,
		},
		{
			description: "uses the provided `halfTowardsZero` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfTowardsZeroBigInt,
		},
		{
			description: "uses the provided `halfAwayFromZero` divide function A",
			input:       dinero.NewBigDineroWithScale(33, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(6, abc, 1),
			divide:      divide.HalfAwayFromZeroBigInt,
		},
		{
			description: "uses the provided `halfAwayFromZero` divide function B",
			input:       dinero.NewBigDineroWithScale(39, abc, 2),
			scale:       big.NewInt(1),
			expect:      dinero.NewBigDineroWithScale(7, abc, 1),
			divide:      divide.HalfAwayFromZeroBigInt,
		},
	}

	for _, tc := range tests {
		var got dinero.Dinero[*big.Int]
		var err error

		if tc.divide != nil {
			got, err = tc.input.TransformScale(tc.scale, tc.divide)
		} else {
			got, err = tc.input.TransformScale(tc.scale, divide.DownBigInt)
		}

		if err != nil {
			t.Errorf("error: %v, %e", tc.expect, err)
		}

		if !reflect.DeepEqual(tc.expect, got) {
			t.Errorf("expected: %v, got: %v", tc.expect, got)
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
