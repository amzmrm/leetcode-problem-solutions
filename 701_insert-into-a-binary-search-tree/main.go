package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root // 把val放到root的left或者right后，还是返回了root，这就确保了树原本的结构没有被改变
}
