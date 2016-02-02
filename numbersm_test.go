package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestStateMachine(t *testing.T) {
	sm := NumberSM{InitialState, ""}
	for _, c := range "1222.020 2" {
		if !sm.Feed(string(c)) {
			break
		}
	}
	AssertEqual(t, sm.IsNumber(), true)
	AssertEqual(t, sm.Records(), "1222.020")
}
