package dinero_test

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/DustinJSilk/dinero.go/currency"
	"github.com/DustinJSilk/dinero.go/dinero"
)

var BigUSD = currency.Currency[*big.Int]{
	Code:     "USD",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(2),
}

var BigEUR = currency.Currency[*big.Int]{
	Code:     "EUR",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(2),
}

var BigMGA = currency.Currency[*big.Int]{
	Code:     "MGA",
	Base:     big.NewInt(5),
	Exponent: big.NewInt(1),
}

var BigMRU = currency.Currency[*big.Int]{
	Code:     "MRU",
	Base:     big.NewInt(5),
	Exponent: big.NewInt(1),
}

var BigIQD = currency.Currency[*big.Int]{
	Code:     "IQD",
	Base:     big.NewInt(10),
	Exponent: big.NewInt(3),
}

type container struct {
	Dinero dinero.Dinero[int] `json:"dinero"`
}

func TestJSON(t *testing.T) {
	d := dinero.NewDinero(1234, currency.USD)
	c := container{d}

	b, err := json.Marshal(c)
	if err != nil {
		t.Errorf("error marshalling: %v", err)
	}

	// Expect the correct json output
	expected := `{"dinero":{"amount":1234,"currency":{"code":"USD","base":10,"exponent":2},"scale":2}}`
	if string(b) != expected {
		t.Errorf("expected json: %v, got: %v", expected, string(b))
	}

	out := container{}
	err = json.Unmarshal(b, &out)
	if err != nil {
		t.Errorf("error unmarshalling: %v", err)
	}

	// Expect calculations to work after unmarshalling with a nil calculator
	multiplied := out.Dinero.Multiply(2)
	withCalculator := multiplied.WithCalculator(dinero.IntCalculator)
	multipledExpect := dinero.NewDinero(2468, currency.USD)
	if !reflect.DeepEqual(withCalculator, multipledExpect) {
		t.Errorf("expected calculations to work: %v, got: %v", multipledExpect, withCalculator)
	}

	// Expect the unmarshalled value to equal the original input value
	if !reflect.DeepEqual(container{out.Dinero.WithCalculator(dinero.IntCalculator)}, c) {
		t.Errorf("expected unmarshalled: %v, got: %v", c, out)
	}
}
