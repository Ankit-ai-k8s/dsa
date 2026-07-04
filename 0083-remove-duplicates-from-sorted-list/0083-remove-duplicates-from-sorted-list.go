/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	prev := head
	for cur != nil {
		cv := cur.Val
		pv := prev.Val
		if cv != pv {
			prev.Next = cur
			prev = cur
		}
		cur = cur.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return head
}