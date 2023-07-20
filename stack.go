package gtree

import (
	"container/list"
	"errors"
)

var errNilStack = errors.New("nil stack")

type stack struct {
	nodes *list.List
}

func newStack() *stack {
	return &stack{nodes: list.New()}
}

func (s *stack) push(n *Node) *stack {
	s.nodes.PushBack(n)
	return s
}

func (s *stack) pop() *Node {
	tmp := s.nodes.Back()
	if tmp == nil {
		return nil
	}

	n := s.nodes.Remove(tmp)
	return n.(*Node)
}

func (s *stack) size() int {
	return s.nodes.Len()
}

// depth-first search
func (s *stack) dfs(current *Node) {
	size := s.size()
	for i := 0; i < size; i++ {
		parent := s.pop()

		// for same name on the same hierarchy
		if child := parent.findChildByText(current.name); child != nil {
			if !child.isDirectlyUnder(parent) {
				continue
			}
			s.push(parent).push(child)
		} else {
			if !current.isDirectlyUnder(parent) {
				continue
			}
			parent.addChild(current)
			current.setParent(parent)
			s.push(parent).push(current)
		}
		return
	}
}
