/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left)
	r := dfs(root.Right)

	return Max(l, r) + 1
}

func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}