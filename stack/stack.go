package stack

import "fmt"

type stack struct {
	items []int
}

func StackOps() {
	stack := stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.print()
	// stack.peek()
	stack.pop()
	// stack.pop()
	// stack.pop()
	stack.print()
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) peek() {
	val := s.items[len(s.items)-1]
	fmt.Println("top: ", val)
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) pop() {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return
	}
	fmt.Println("popped: ", s.items[len(s.items)-1])
	s.items = s.items[:len(s.items)-1]
}

func (s *stack) print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}
