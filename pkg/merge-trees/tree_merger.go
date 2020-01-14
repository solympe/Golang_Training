package merge_trees

// TreeMerger ...
type TreeMerger interface {
	MergeTrees(t1 TreeMerger, t2 TreeMerger) TreeMerger
	GetLeft() TreeMerger
	GetRight() TreeMerger
	GetVal() int
	setVal(val int)
	setLeft(node TreeMerger)
	setRight(node TreeMerger)
}
