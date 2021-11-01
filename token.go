package calc

import "fmt"

type TokenKind int

const (
	_ TokenKind = iota
	TkReserved
	TkNum
	TkIdent
	TkEof
)

var kinds = [...]string{
	TkReserved: "TkReserved",
	TkNum:      "TkNum",
	TkEof:      "TkEof",
}

func (tokKind TokenKind) String() string {
	return kinds[tokKind]
}

func NewToken(kind TokenKind, pos [2]int, lit string, val int) Token {
	return Token{
		kind: kind,
		pos:  pos,
		val:  val,
		lit:  lit,
	}
}

type Token struct {
	kind TokenKind // 種類
	pos  [2]int    // オリジナル文字の位置
	val  int       // Numだった場合に数値が入る
	lit  string    // オリジナル文字
}

func (tok Token) String() string {
	return fmt.Sprintf("Token{ kind: %v, valInt: %v, lit: %v, pos: %v-%v }", tok.kind.String(), tok.val, tok.lit, tok.pos[0], tok.pos[1])
}
