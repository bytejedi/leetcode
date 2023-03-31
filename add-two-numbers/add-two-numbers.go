package add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := new(ListNode)
	current := first
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + current.Val
		v := sum - 10
		if v >= 0 {
			current.Val = v
			current.Next = &ListNode{Val: 1}
			current = current.Next
		} else {
			current.Val = sum
			if l1 != nil || l2 != nil {
				current.Next = &ListNode{}
				current = current.Next
			}
		}
	}
	return first
}
