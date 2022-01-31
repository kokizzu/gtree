// Package gtree provides tree-structured output.
package gtree

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

var (
	idxCounter = newCounter()
)

var (
	// ErrNilNode is returned if the argument *gtree.Node of OutputProgrammably function is nill.
	ErrNilNode = errors.New("nil node")
	// ErrNotRoot is returned if the argument *gtree.Node of OutputProgrammably function is not root of the tree.
	ErrNotRoot = errors.New("not root node")
)

// OutputProgrammably outputs tree to w.
// This function requires node generated by NewRoot function.
func OutputProgrammably(w io.Writer, root *Node, optFns ...OptFn) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}

	idxCounter.reset()

	tree := newTree(conf.encode, conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun, conf.fileExtensions)
	tree.addRoot(root)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
}

// MkdirProgrammably makes directories.
// This function requires node generated by NewRoot function.
func MkdirProgrammably(root *Node, optFns ...OptFn) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}

	idxCounter.reset()

	tree := newTree(conf.encode, conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun, conf.fileExtensions)
	tree.addRoot(root)

	if conf.dryrun {
		// when detect invalid node name, return error. process end.
		// when detected no invalid node name, output tree. process end.
		if err := tree.grow(); err != nil {
			return err
		}
		return tree.spread(os.Stdout)
	}

	// 微妙?
	tree.setDryRun(true)
	// when detect invalid node name, return error. process end.
	// when detected no invalid node name, no output tree. process continue.
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}

// NewRoot creates a starting node for building tree.
func NewRoot(text string) *Node {
	return newNode(text, rootHierarchyNum, idxCounter.next())
}

// Add adds a node and returns an instance of it.
// If a node with the same text already exists in the same hierarchy of the tree, that node will be returned.
func (parent *Node) Add(text string) *Node {
	for _, child := range parent.Children {
		if text == child.Name {
			return child
		}
	}

	current := newNode(text, parent.hierarchy+1, idxCounter.next())
	current.setParent(parent)
	parent.addChild(current)
	return current
}
