package gcalc

type ParsePhase int64

const (
	ParseInitail ParsePhase = iota + 1
	PrimaryExp
	Term
	Exp
)

type ParseState struct {
	State       ParsePhase
	Transitions map[[2]TokenKind]*ParseState
}

func NewParseState(ps ParsePhase) *ParseState {
	return &ParseState{ps, map[[2]TokenKind]*ParseState{}}
}

type ParseSM struct {
	CurrentState *ParseState
}

func (psm *ParseSM) Feed(firstT, secodeT *Token) bool {
	if nextState, ok := psm.CurrentState.Transitions[tokenQueue(firstT.kind, secodeT.kind)]; ok {
		psm.CurrentState = nextState
		return true
	}

	return false
}

var ParseInitailState = NewParseState(ParseInitail)
var PrimaryExpState = NewParseState(PrimaryExp)
var TermState = NewParseState(Term)
var ExpState = NewParseState(Exp)

func tokenQueue(a, b TokenKind) [2]TokenKind {
	return [2]TokenKind{a, b}
}

/* Simple caculator BNF graph

Exp ------------------------------------------------->
      |                                       |
      |<----------- Term Add/Sub -------------|


Term -----PrimaryExp--------------------------------->
              |                               |
              |<---- PrimaryExp Mul/Div <-----|


PrimaryExp -----------Number------------------------->

*/

func init() {
	ParseInitailState.Transitions[tokenQueue(NumberToken, NumberToken)] = PrimaryExpState

	PrimaryExpState.Transitions[tokenQueue(AddOperatorToken, NumberToken)] = ExpState
	PrimaryExpState.Transitions[tokenQueue(SubOperatorToken, NumberToken)] = ExpState
	PrimaryExpState.Transitions[tokenQueue(MulOperatorToken, NumberToken)] = TermState
	PrimaryExpState.Transitions[tokenQueue(DivOperatorToken, NumberToken)] = TermState

	TermState.Transitions[tokenQueue(AddOperatorToken, NumberToken)] = ExpState
	TermState.Transitions[tokenQueue(SubOperatorToken, NumberToken)] = ExpState
	TermState.Transitions[tokenQueue(MulOperatorToken, NumberToken)] = TermState
	TermState.Transitions[tokenQueue(DivOperatorToken, NumberToken)] = TermState

	ExpState.Transitions[tokenQueue(AddOperatorToken, NumberToken)] = ExpState
	ExpState.Transitions[tokenQueue(SubOperatorToken, NumberToken)] = ExpState
	ExpState.Transitions[tokenQueue(MulOperatorToken, NumberToken)] = ExpState
	ExpState.Transitions[tokenQueue(DivOperatorToken, NumberToken)] = ExpState
}
