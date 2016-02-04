package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := NewCalculator("(1.00 + 2) * 3-(1.0+ 2)* 3 ")
	t.Log(string(calc.exp))
	// t.Log(calc.nextToken())
	AssertEqual(t, calc.Calculate(), float64(0))
}
