package calc

// expr       = equality
// equality   = relational ("==" equality | "!=" equality)*
// relational = add (">" add | ">=" add | "<" add | "<=" add)*
// add        = mul ("+" mul | "-" mul)*
// mul        = unary ("*" unary | "/" unary)*
// unary      = ("+" | "-")? primary
// primary    = num | "(" expr ")"

func NewParser(tokens []Token) Parser {
	return Parser{
		tokens: tokens,
		pos:    0,
	}
}

type Parser struct {
	tokens []Token
	pos    int
}

func (ps *Parser) curt() Token {
	return ps.tokens[ps.pos]
}
func (ps *Parser) next() Token {
	return ps.tokens[ps.pos+1]
}

func (ps *Parser) goNext() {
	ps.pos++
}

func (ps *Parser) isEof() bool {
	return ps.curt().kind == TkEof
}

func (ps *Parser) consume(op string) bool {
	if ps.curt().kind != TkReserved || ps.curt().lit != op {
		return false
	}
	ps.goNext()
	return true
}

func (ps *Parser) expect(op string) bool {
	if ps.curt().lit == op {
		ps.goNext()
		return true
	} else {
		return false
	}
}

func (ps *Parser) expectNum() int {
	if ps.curt().kind != TkNum {
		panic("unexpected")
	}

	val := ps.curt().val
	ps.goNext()
	return val
}

func (ps *Parser) Expr() Node {
	return ps.Equality()
}

func (ps *Parser) Equality() Node {
	node := ps.Relational()

	if ps.consume("==") {
		node = NewNode(NdEqual, node, ps.Relational())
	} else if ps.consume("!=") {
		node = NewNode(NdNotEqual, node, ps.Relational())
	}

	return node
}

func (ps *Parser) Relational() Node {
	node := ps.Add()
	if ps.consume("<") {
		node = NewNode(NdLt, node, ps.Add())
	} else if ps.consume("<=") {
		node = NewNode(NdLte, node, ps.Add())
	} else if ps.consume(">") {
		node = NewNode(NdGt, node, ps.Add())
	} else if ps.consume(">=") {
		node = NewNode(NdGte, node, ps.Add())
	}

	return node
}

func (ps *Parser) Add() Node {
	node := ps.Mul()

	for !ps.isEof() {
		if ps.consume("+") {
			node = NewNode(NdAdd, node, ps.Mul())
		} else if ps.consume("-") {
			node = NewNode(NdSub, node, ps.Mul())
		} else {
			break
		}
	}
	return node
}

func (ps *Parser) Mul() Node {
	node := ps.Unary()

	for !ps.isEof() {
		if ps.consume("*") {
			node = NewNode(NdMul, node, ps.Unary())
		} else if ps.consume("/") {
			node = NewNode(NdDiv, node, ps.Unary())
		} else {
			break
		}
	}

	return node
}

func (ps *Parser) Unary() Node {
	if ps.consume("+") {
		return ps.Primary()
	} else if ps.consume("-") {
		return NewNode(NdSub, NewNumNode(0), ps.Primary())
	}
	return ps.Primary()
}

func (ps *Parser) Primary() Node {
	if ps.consume("(") {
		node := ps.Expr()
		ps.expect(")")
		return node
	}
	return NewNumNode(ps.expectNum())
}
