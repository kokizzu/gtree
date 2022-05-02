package gtree

// 関心事は各ノードの枝を組み立てること
type grower interface {
	grow([]*Node) error
	enableValidation()
}

func newGrower(
	encode encode,
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) grower {
	if encode != encodeDefault {
		return &noopGrower{}
	}
	return &defaultGrower{
		lastNodeFormat:        lastNodeFormat,
		intermedialNodeFormat: intermedialNodeFormat,
		enabledValidation:     enabledValidation,
	}
}

type branchFormat struct {
	directly, indirectly string
}

type defaultGrower struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat
	enabledValidation     bool
}

func (dg *defaultGrower) grow(roots []*Node) error {
	for _, root := range roots {
		if err := dg.assemble(root); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrower) assemble(current *Node) error {
	if err := dg.assembleBranch(current); err != nil {
		return err
	}

	for _, child := range current.Children {
		if err := dg.assemble(child); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrower) assembleBranch(current *Node) error {
	if current.isRoot() {
		return nil
	}

	dg.assembleBranchDirectly(current)

	// go back to the root to form a brnch
	tmpParent := current.parent
	for {
		if !tmpParent.isRoot() {
			dg.assembleBranchIndirectly(current, tmpParent)
			tmpParent = tmpParent.parent
			continue
		}

		dg.assembleBranchFinally(current, tmpParent)
		if !dg.enabledValidation {
			return nil
		}
		if err := current.validatePath(); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (dg *defaultGrower) assembleBranchDirectly(current *Node) {
	current.setPath(current.Name)

	if current.isLastOfHierarchy() {
		current.setBranch(current.branch(), dg.lastNodeFormat.directly)
	} else {
		current.setBranch(current.branch(), dg.intermedialNodeFormat.directly)
	}
}

func (dg *defaultGrower) assembleBranchIndirectly(current, parent *Node) {
	current.setPath(parent.Name, current.path())

	if parent.isLastOfHierarchy() {
		current.setBranch(dg.lastNodeFormat.indirectly, current.branch())
	} else {
		current.setBranch(dg.intermedialNodeFormat.indirectly, current.branch())
	}
}

func (*defaultGrower) assembleBranchFinally(current, root *Node) {
	current.setPath(root.path(), current.path())
}

func (dg *defaultGrower) enableValidation() {
	dg.enabledValidation = true
}

type noopGrower struct{}

func (*noopGrower) grow(_ []*Node) error { return nil }

func (*noopGrower) enableValidation() {}
