package calc

import (
	"fmt"
	"strconv"
	"unicode"
)

func NewTokenizer(text string) Tokenizer {
	return Tokenizer{
		pos:   0,
		runes: []rune(text),
	}
}

type Tokenizer struct {
	pos   int
	runes []rune
}

func (tk *Tokenizer) curt() rune {
	return tk.runes[tk.pos]
}
func (tk *Tokenizer) next() rune {
	return tk.runes[tk.pos+1]
}

func (tk *Tokenizer) goNext() {
	tk.pos++
}

func (tk *Tokenizer) isEof() bool {
	return tk.pos >= len(tk.runes)
}

func (tk *Tokenizer) consumeNumber() (Token, error) {
	numStr := ""
	s := tk.pos
	for !tk.isEof() {
		r := tk.curt()
		if unicode.IsDigit(r) {
			numStr += string(r)
			tk.goNext()
		} else {
			break
		}
	}
	e := tk.pos
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return Token{}, err
	}

	return NewToken(TkNum, [2]int{s, e}, numStr, num), nil
}

func (tk *Tokenizer) IsOperator(r rune) bool {
	for _, opr := range []rune("+-*/!=<>") {
		if opr == r {
			return true
		}
	}
	return false
}

func (tk *Tokenizer) consumeOperator() Token {
	opr := ""
	s := tk.pos
	for !tk.isEof() {
		r := tk.curt()
		if tk.IsOperator(r) {
			opr += string(r)
		} else {
			break
		}
		tk.goNext()
	}
	e := tk.pos

	return NewToken(TkReserved, [2]int{s, e}, opr, 0)
}

func (tk *Tokenizer) IsParen(r rune) bool {
	for _, opr := range []rune("()") {
		if opr == r {
			return true
		}
	}
	return false
}

func (tk *Tokenizer) consumeParen() Token {
	s := tk.pos
	r := tk.curt()
	tk.goNext()
	e := tk.pos
	return NewToken(TkReserved, [2]int{s, e}, string(r), 0)
}

func (tk *Tokenizer) Tokenize() ([]Token, error) {
	var tokens []Token

	for !tk.isEof() {
		r := tk.curt()
		if unicode.IsDigit(r) {
			tok, err := tk.consumeNumber()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tok)
		} else if tk.IsOperator(r) {
			tokens = append(tokens, tk.consumeOperator())
		} else if tk.IsParen(r) {
			tokens = append(tokens, tk.consumeParen())
		} else if unicode.IsSpace(r) {
			tk.goNext()
		} else {
			return nil, fmt.Errorf("syntax error : %v", string(r))
		}
	}

	tokens = append(tokens, NewToken(TkEof, [2]int{tk.pos, tk.pos + 1}, "", 0))

	return tokens, nil
}
