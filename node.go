package calc

type NodeKind int

const (
	_          NodeKind = iota
	NdAdd               // +
	NdSub               // -
	NdMul               // *
	NdDiv               // /
	NdEqual             // ==
	NdNotEqual          // !=
	NdLt                // <
	NdLte               // <=
	NdGt                // >
	NdGte               // >=
	NdOr                // ||
	NdAnd               // &&
	NdAssign            // =
	NdStmt              // ;

	NdNum // 整数
	NdIdent
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
		kind:   NdNum,
		valInt: val,
	}
}

func NewIdentNode(val string) Node {
	return Node{
		kind:      NdIdent,
		valString: val,
	}
}

type Node struct {
	kind      NodeKind
	lhs       *Node
	rhs       *Node
	valInt    int
	valString string
	valFloat  float64
	valBool   bool
}
