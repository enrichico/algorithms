//Package stringstack contains implementation of a stack of strings
package stringstack

import "fmt"

type node struct {
	item string
	next *node
}

//Type Stack provides the implementation of a stack of strings as a linked list
type Stack struct {
	first *node
}

func New() *Stack {

	return new(Stack)
}

func (s *Stack) IsEmpty() bool {

	return (s.first == nil)
}

func (s *Stack) Push(item string) {

	oldFirst := s.first
	s.first = new(node)
	s.first.item = item
	s.first.next = oldFirst
}

func (s *Stack) Pop() string {

	item := s.first.item
	s.first = s.first.next
	return item
}

func (s *Stack) Print() {

	for e := s.first; e != nil; e = e.next {
		// do something with e.item
		fmt.Printf("%s\n", e.item)
	}

}
