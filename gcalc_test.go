package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestCalculator(t *testing.T) {
	result, err := Calculate("  1 + 2* 4 - 1 -8 /1")
	AssertNilError(t, err)
	AssertEqual(t, result, 0.0)
}
