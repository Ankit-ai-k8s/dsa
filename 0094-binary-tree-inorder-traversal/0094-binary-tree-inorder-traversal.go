/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	op := &[]int{}
	dfs(root, op)
	return *op
}

func dfs(root *TreeNode, op *[]int) {
    if root == nil {
		return
	}

    dfs(root.Left, op)
    *op = append(*op, root.Val)
    dfs(root.Right, op)
}