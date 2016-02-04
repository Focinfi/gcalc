package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := NewCalculator(" 1*(-1)* (5.00 + (-2)) * 3-(1.0+ 2)* 3 *(-1)")
	AssertEqual(t, calc.Calculate(), float64(0))
}
