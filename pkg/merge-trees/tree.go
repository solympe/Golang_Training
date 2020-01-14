package merge_trees

type tree struct {
	val   int
	left  TreeMerger
	right TreeMerger
}

// GetLeft ...
func (t *tree) GetLeft() TreeMerger {
	return t.left
}

// GetRight ...
func (t *tree) GetRight() TreeMerger {
	return t.right
}

// GetVal ...
func (t *tree) GetVal() int {
	return t.val
}

func (t *tree) setLeft(node TreeMerger) {
	t.left = node
}

func (t *tree) setRight(node TreeMerger) {
	t.right = node
}

func (t *tree) setVal(val int) {
	t.val = val
}

// MergeTrees ...
func (t *tree) MergeTrees(t1 TreeMerger, t2 TreeMerger) TreeMerger {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}
	t1.setVal(t2.GetVal() + t1.GetVal())
	t1.setLeft(t1.MergeTrees(t1.GetLeft(), t2.GetLeft()))
	t1.setRight(t1.MergeTrees(t1.GetRight(), t2.GetRight()))

	return t1
}

// NewTreeMerger ...
func NewTreeMerger(val int, left TreeMerger, right TreeMerger) TreeMerger {
	return &tree{
		val:   val,
		left:  left,
		right: right,
	}
}
