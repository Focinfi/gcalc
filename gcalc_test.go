package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestCaculator(t *testing.T) {
	result, err := Caculate("  1 + 2* 4 - 1 -8 /1")
	AssertNilError(t, err)
	AssertEqual(t, result, 0.0)
}
