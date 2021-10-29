package calc

type NodeKind int

const (
	_          NodeKind = iota
	NdAdd               // +
	NdSub               // -
	NdMul               // *
	NdDiv               // /
	NdLt                // <
	NdLte               // <=
	NdGt                // >
	NdGte               // >=
	NdEqual             // ==
	NdNotEqual          // !=
	NdNum               // 整数
)

func NewNode(kind NodeKind, lhs Node, rhs Node) Node {
	return Node{
		kind: kind,
		lhs:  &lhs,
		rhs:  &rhs,
	}
}

func NewNumNode(val int) Node {
	return Node{
		kind: NdNum,
		val:  val,
	}
}

type Node struct {
	kind NodeKind
	lhs  *Node
	rhs  *Node
	val  int
}
