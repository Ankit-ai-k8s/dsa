/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slw := head
    fst := head

    for fst != nil && fst.Next != nil {
        slw = slw.Next
        fst = fst.Next.Next
    }

    return slw
}