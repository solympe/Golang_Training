package merge_trees

// TreeMerger ...
type TreeMerger interface {
	GetLeft() TreeMerger
	GetRight() TreeMerger
	GetVal() int
	setVal(val int)
	setLeft(node TreeMerger)
	setRight(node TreeMerger)
}
