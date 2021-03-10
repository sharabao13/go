package main

// 线段树
type SegTreeNode struct {
	start int
	end   int
	min   int
	left  *SegTreeNode
	right *SegTreeNode
}

// 建树
func build(nums []int, start, end int) *SegTreeNode {
	if start > end {
		return nil
	}
	root := &SegTreeNode{start, end, start, nil, nil}
	if start == end {
		return root
	}
	mid := (start + end) >> 1
	root.left = build(nums, start, mid)
	root.right = build(nums, mid+1, end)
	root.min = root.left.min
	if nums[root.min] > nums[root.right.min] {
		root.min = root.right.min
	}
	return root
}

func query(nums []int, root *SegTreeNode, start, end int) int {
	if root == nil || end < root.start || start > root.end {
		return -1
	}
	if start <= root.start && end >= root.end {
		return root.min
	}
	lmin := query(nums, root.left, start, end)
	rmin := query(nums, root.right, start, end)
	if lmin == -1 {
		return rmin
	}
	if rmin == -1 {
		return lmin
	}
	if nums[lmin] > nums[rmin] {
		return rmin
	}
	return lmin
}

// 分治一下
func helper(heights []int, root *SegTreeNode, start, end int) int {
	if start > end {
		return -1
	}
	idxMin := query(heights, root, start, end)
	// fmt.Println(idxMin, heights[idxMin], start, end)
	return max(heights[idxMin]*(end-start+1), max(helper(heights, root, start, idxMin-1), helper(heights, root, idxMin+1, end)))
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	root := build(heights, 0, len(heights)-1)
	res := helper(heights, root, 0, len(heights)-1)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
