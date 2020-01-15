package main

import (
	"fmt"
)

type Node struct {
	value byte
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
func (s *Stack) push(value byte) {
	if s.top == nil {
		s.top = &Node{value, nil}
	} else {
		next := s.top
		s.top = &Node{value, next}
	}
	s.h++
}
func (s *Stack) pop() byte {
	if s.top == nil {
		return '#'
	}
	value := s.top.value
	s.top = s.top.next
	s.h--
	return value
}

// 判断是否是有效括号
func isValid(s string) bool {
	stack := Init()
	for i := 0; i < len(s); i++ {
		pre := stack.pop()
		if pre == '#' {
			stack.push(s[i])
		} else {

			if s[i]-pre == '2'-'0' || s[i]-pre == '1'-'0' {
				continue
			} else {
				stack.push(pre)
				stack.push(s[i])
			}
		}
	}
	if stack.pop() == '#' {
		return true
	}
	return false
}
func main() {
	fmt.Println(isValid("(([]){})"))
}
