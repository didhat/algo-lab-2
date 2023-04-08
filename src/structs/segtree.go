package structs

type SegTreeNode struct {
	left  *SegTreeNode
	right *SegTreeNode
	sum   int
}

func NewSegTreeNode(left, right *SegTreeNode, sum int) SegTreeNode {
	return SegTreeNode{left: left, right: right, sum: sum}
}

func NewEmptySegTreeNode() SegTreeNode {
	return SegTreeNode{left: nil, right: nil, sum: 0}
}

func AddToSegTree(root SegTreeNode, left, right, rangeStart, rangeEnd, value int) SegTreeNode {
	if left >= rangeEnd || right <= rangeStart {
		return root
	}

	if rangeStart <= left && right <= rangeEnd {
		newRoot := NewSegTreeNode(root.left, root.right, root.sum)
		newRoot.sum += value
		return newRoot
	}

	mid := (left + right) / 2
	newRoot := NewSegTreeNode(root.left, root.right, root.sum)

	if newRoot.left == nil {
		_left := NewEmptySegTreeNode()
		newRoot.left = &_left
	}

	newLeft := AddToSegTree(*newRoot.left, left, mid, rangeStart, rangeEnd, value)
	newRoot.left = &newLeft

	if newRoot.right == nil {
		_right := NewEmptySegTreeNode()
		newRoot.right = &_right
	}
	newRight := AddToSegTree(*newRoot.right, mid, right, rangeStart, rangeEnd, value)
	newRoot.right = &newRight

	return newRoot
}

func GetSum(root SegTreeNode, left, right, targetZippedX int) int {
	if right-left == 1 {
		return root.sum
	}

	mid := (left + right) / 2

	if targetZippedX < mid {
		if root.left == nil {
			return root.sum
		}
		return root.sum + GetSum(*root.left, left, mid, targetZippedX)
	} else {
		if root.right == nil {
			return root.sum
		}
		return root.sum + GetSum(*root.right, mid, right, targetZippedX)
	}

}
