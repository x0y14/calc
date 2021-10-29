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
	rdi   int
	rax   int
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

func (c *Calculator) Is(expression string) bool {
	c.Calculate(expression)
	if c.rax == 1 {
		return true
	}
	return false
}

func (c *Calculator) push(val int) {
	// now:
	// [ 0 0 0 0 0 ]
	//           ^
	//           sp
	// pushed:
	// [ 0 0 0 val 0 ]
	//         ^
	//         sp

	c.sp--
	c.stack[c.sp] = val
}
func (c *Calculator) pop() int {
	// now:
	// [ 0 0 0 val 0 ]
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

func (c *Calculator) traverse(node Node) {
	if node.kind == NdNum {
		c.push(node.val)
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
		// [ 0, 0, 1, 2, 3 ]
		//         ^
		c.rdi = c.pop() // 1
		// [ 0, 0, 0, 2, 3 ]
		//            ^
		c.rax = c.pop() // 2
		// [ 0, 0, 0, 0, 3 ]
		//               ^
		if c.rdi == c.rax {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}
		c.rax = c.al
	case NdNotEqual:
		c.rdi = c.pop()
		c.rax = c.pop()
		if c.rdi != c.rax {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}
		c.rax = c.al
	case NdLt:
		c.rdi = c.pop()
		c.rax = c.pop()

		if c.rax < c.rdi {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}

		c.rax = c.al
	case NdLte:
		c.rdi = c.pop()
		c.rax = c.pop()

		if c.rax <= c.rdi {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}

		c.rax = c.al
	case NdGt:
		c.rdi = c.pop()
		c.rax = c.pop()

		if c.rax > c.rdi {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}

		c.rax = c.al
	case NdGte:
		c.rdi = c.pop()
		c.rax = c.pop()

		if c.rax >= c.rdi {
			c.al = 1
			//c.push(1)
		} else {
			c.al = 0
			//c.push(0)
		}

		c.rax = c.al
	}
	return
}
