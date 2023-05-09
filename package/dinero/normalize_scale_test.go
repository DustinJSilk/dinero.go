package dinero_test

import (
	"reflect"
	"testing"

	"dinero.go/package/currency"
	"dinero.go/package/dinero"
)

func TestNormalizeScale(t *testing.T) {
	type test struct {
		a       dinero.Dinero[int]
		b       dinero.Dinero[int]
		expectA dinero.Dinero[int]
		expectB dinero.Dinero[int]
	}

	tests := []test{
		{
			a:       dinero.NewDineroWithScale(100, currency.USD, 2),
			b:       dinero.NewDineroWithScale(1000, currency.USD, 3),
			expectA: dinero.NewDineroWithScale(1000, currency.USD, 3),
			expectB: dinero.NewDineroWithScale(1000, currency.USD, 3),
		},
		{
			a:       dinero.NewDinero(500, currency.USD),
			b:       dinero.NewDineroWithScale(1000, currency.USD, 3),
			expectA: dinero.NewDineroWithScale(5000, currency.USD, 3),
			expectB: dinero.NewDineroWithScale(1000, currency.USD, 3),
		},
	}

	for _, tc := range tests {
		got, err := dinero.NormalizeScale(tc.a, tc.b)
		if err != nil {
			t.Fatalf("error: %v, %v, %e", tc.a, tc.b, err)
		}

		if !reflect.DeepEqual(tc.expectA, got[0]) {
			t.Fatalf("expected a: %v, got: %v", tc.expectA, got[0])
		}
		if !reflect.DeepEqual(tc.expectB, got[1]) {
			t.Fatalf("expected b: %v, got: %v", tc.expectB, got[1])
		}
	}
}

func BenchmarkNormalizeScale(b *testing.B) {
	da := dinero.NewDineroWithScale(100, currency.USD, 2)
	db := dinero.NewDineroWithScale(1000, currency.USD, 2)
	dc := dinero.NewDineroWithScale(10000, currency.USD, 3)

	for i := 0; i < b.N; i++ {
		_, err := dinero.NormalizeScale(da, db, dc)
		if err != nil {
			b.Fatalf("error: %e", err)
		}
	}
}
