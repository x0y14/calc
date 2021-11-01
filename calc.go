package calc

func NewCalculator(stackSize int) Calculator {
	return Calculator{
		stack: make([]int, stackSize),
		sp:    stackSize - 1,
	}
}

type Calculator struct {
	stack []int
	sp    int
	rax   int
	rdi   int
	al    int
}

func (c *Calculator) Calculate(expression string) int {
	tk := NewTokenizer(expression)
	tokens, err := tk.Tokenize()
	if err != nil {
		panic(err)
	}

	ps := NewParser(tokens)
	node := ps.Expr()
	c.traverse(node)
	return c.pop()
}

func (c *Calculator) push(val int) {
	// now:
	// [ 0 0 0 0 0 ]
	//           ^
	//           sp
	// pushed:
	// [ 0 0 0 valInt 0 ]
	//         ^
	//         sp

	c.sp--
	c.stack[c.sp] = val
}
func (c *Calculator) pop() int {
	// now:
	// [ 0 0 0 valInt 0 ]
	//         ^
	//         sp
	// popped:
	// [ 0 0 0 0 0 ]
	//           ^
	//           sp
	val := c.stack[c.sp]
	c.stack[c.sp] = 0
	c.sp++
	return val
}

func (c *Calculator) pop2r() {
	c.rax = c.pop()
	c.rdi = c.pop()
}

func (c *Calculator) Is(text string) bool {
	c.Calculate(text)
	if c.rax == 1 {
		return true
	}
	return false
}

func (c *Calculator) traverse(node Node) {
	if node.kind == NdNum {
		c.push(node.valInt)
		return
	}

	c.traverse(*node.lhs)
	c.traverse(*node.rhs)

	switch node.kind {
	case NdAdd:
		a := c.pop()
		b := c.pop()
		c.push(a + b)
	case NdSub:
		a := c.pop()
		b := c.pop()
		c.push(b - a)
	case NdMul:
		a := c.pop()
		b := c.pop()
		c.push(a * b)
	case NdDiv:
		a := c.pop()
		b := c.pop()
		c.push(b / a)

	case NdEqual:
		c.pop2r()
		if c.rdi == c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	case NdNotEqual:
		c.pop2r()
		if c.rdi != c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	case NdLt:
		c.pop2r()
		if c.rdi < c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	case NdLte:
		c.pop2r()
		if c.rdi <= c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	case NdGt:
		c.pop2r()
		if c.rdi > c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	case NdGte:
		c.pop2r()
		if c.rdi >= c.rax {
			c.al = 1
		} else {
			c.al = 0
		}
		c.rax = c.al
	}
	return
}
