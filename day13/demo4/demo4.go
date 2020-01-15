package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
	h   int
}

func Init() *Stack {
	return &Stack{
		nil,
		0,
	}
}
func (s *Stack) push(value int) {
	if s.top == nil {
		s.top = &Node{value, nil}
	} else {
		next := s.top
		s.top = &Node{value, next}
	}
	s.h++
}
func (s *Stack) pop() int {
	if s.top == nil {
		return '#'
	}
	value := s.top.value
	s.top = s.top.next
	s.h--
	return value
}
func (s *Stack) peek() int {
	if s.top == nil {
		return '#'
	}
	value := s.top.value
	return value
}
func longestValidParentheses(s string) int {
	max := 0
	sa := Init()
	sa.push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sa.push(i)
		} else {

			sa.pop()
			temp := sa.peek()

			if sa.h == 0 {
				sa.push(i)
			} else {
				if max < (i - temp) {
					max = i - temp
				}
			}
		}
	}
	return max
}
func main() {
	fmt.Println(longestValidParentheses("())((())"))
}
