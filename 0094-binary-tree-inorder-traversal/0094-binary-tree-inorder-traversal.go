/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	return dfs(root, []int{})
}

func dfs(root *TreeNode, op []int) []int {
	if root == nil {
		return op
	}

	op = dfs(root.Left, op)
	op = append(op, root.Val)
	op = dfs(root.Right, op)

	return op
}