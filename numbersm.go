package gcalc

import (
	"strconv"
)

type FloatPartState int64

const (
	FloatInitial FloatPartState = 1 << (iota + 1)
	FloatInIntPart
	FloatDot
	FloatInFracPart
)

type NumberSM struct {
	CurrentState *State
	records      string
}

func (sm *NumberSM) Reset() {
	sm.CurrentState = InitialState
	sm.records = ""
}

func (sm *NumberSM) Feed(str string) bool {
	if state, ok := sm.CurrentState.Transitions[str]; ok {
		sm.CurrentState = state
		sm.records += str
		return true
	}

	return false
}

func (sm *NumberSM) IsNumber() bool {
	return sm.CurrentState.FloatPartState != FloatDot && sm.records != ""
}

func (sm *NumberSM) Records() string {
	return sm.records
}

type State struct {
	FloatPartState FloatPartState
	Transitions    map[string]*State
}

func NewStateWithFloatPartState(ls FloatPartState) *State {
	return &State{FloatPartState: ls, Transitions: map[string]*State{}}
}

var InitialState = NewStateWithFloatPartState(FloatInitial)
var InIntPartState = NewStateWithFloatPartState(FloatInIntPart)
var DotState = NewStateWithFloatPartState(FloatDot)
var InFracPartState = NewStateWithFloatPartState(FloatInFracPart)

func init() {
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		InitialState.Transitions[str] = InIntPartState
		InIntPartState.Transitions[str] = InIntPartState
		DotState.Transitions[str] = InFracPartState
		InFracPartState.Transitions[str] = InFracPartState
	}
	InIntPartState.Transitions["."] = DotState
}
