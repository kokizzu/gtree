// Package gtree provides tree-structured output.
package gtree

import (
	"io"

	"github.com/pkg/errors"
)

var (
	// ErrNilNode is returned if the argument *gtree.Node of ExecuteProgrammably function is nill.
	ErrNilNode = errors.New("nil node")
	// ErrNotRoot is returned if the argument *gtree.Node of ExecuteProgrammably function is not root of the tree.
	ErrNotRoot = errors.New("not root node")
)

// ExecuteProgrammably outputs tree to w.
// This function requires node generated by NewRoot function.
func ExecuteProgrammably(w io.Writer, root *Node) error {
	if root == nil {
		return ErrNilNode
	}

	if !root.isRoot() {
		return ErrNotRoot
	}

	tree := &tree{
		roots: []*Node{root},
		lastNodeFormat: lastNodeFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		intermedialNodeFormat: intermedialNodeFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
	}

	tree.grow()
	return tree.expand(w)
}

var programableNodeIdx int

// NewRoot creates a starting node for building tree.
func NewRoot(text string) *Node {
	programableNodeIdx++

	return newNode(text, rootHierarchyNum, programableNodeIdx)
}

// Add adds a node and returns an instance of it.
// If a node with the same text already exists in the same hierarchy of the tree, that node will be returned.
func (parent *Node) Add(text string) *Node {
	for _, child := range parent.children {
		if text == child.text {
			return child
		}
	}

	programableNodeIdx++

	current := newNode(text, parent.hierarchy+1, programableNodeIdx)
	current.parent = parent
	parent.children = append(parent.children, current)
	return current
}
