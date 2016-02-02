package gcalc

import (
	"fmt"
)

type TokenStuck struct {
	tokens []*Token
}

func NewTokenStuck() *TokenStuck {
	return &TokenStuck{[]*Token{}}
}

func (ts *TokenStuck) Len() int {
	return len(ts.tokens)
}

func (ts *TokenStuck) Empty() bool {
	return ts.Len() == 0
}

func (ts *TokenStuck) Clear() {
	ts.tokens = []*Token{}
}

func (ts *TokenStuck) Ret() *Token {
	token := ts.Pop()
	ts.Clear()
	return token
}

func (ts *TokenStuck) Pop() *Token {
	if ts.Empty() {
		return nil
	}

	token := ts.tokens[ts.Len()-1]
	ts.tokens = ts.tokens[:ts.Len()-1]

	return token
}

func (ts *TokenStuck) Push(token *Token) {
	ts.tokens = append(ts.tokens, token)
}

func (ts *TokenStuck) ToSlice() []*Token {
	return ts.tokens
}

func Parse(tokens []Token, stuck *TokenStuck) error {
	length := len(tokens)
	if length == 0 {
		return fmt.Errorf("empty tokens")
	}

	psm := ParseSM{CurrentState: PrimaryExpState}
	currentToken := tokens[0]

	if currentToken.kind != NumberToken {
		return fmt.Errorf("first char[%s] must be a number", currentToken.str)
	} else {
		stuck.Push(&currentToken)
	}

	for i := 1; i < length; i += 2 {
		token := tokens[i]
		if i == length-1 {
			return fmt.Errorf("expression error: ", tokens[i-1], token)
		}

		if !psm.Feed(&token, &tokens[i+1]) {
			return fmt.Errorf("expression error: %v %v", token, tokens[i+1])
		}

		if token.kind == AddOperatorToken || token.kind == SubOperatorToken {
			stuck.Push(&token)
			stuck.Push(&tokens[i+1])
		} else {
			stuck.Push(&Token{value: token.calc(stuck.Pop().value, tokens[i+1].value), kind: NumberToken})
		}
	}

	if stuck.Len() != 1 && psm.CurrentState.State == Exp {
		tokens := stuck.ToSlice()

		if stuck.Len() < 3 || stuck.Len()%2 != 1 {
			return fmt.Errorf("[Parse 1] internal err: result stuck is %v", tokens)
		}

		var sum float64 = tokens[0].value
		for i := 1; i < stuck.Len(); i += 2 {
			if tokens[i].kind == NumberToken || tokens[i+1].kind != NumberToken {
				return fmt.Errorf("[Parse 2] internal err: result stuck is %v", tokens)
			}

			sum = tokens[i].calc(sum, tokens[i+1].value)
		}

		stuck.Clear()
		stuck.Push(&Token{value: sum, kind: NumberToken})
	}

	return nil
}
