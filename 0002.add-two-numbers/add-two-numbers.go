package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	node := &ListNode{}

	lastNode := node
	carry := 0
	for l1 != nil || l2 != nil || carry == 1 {
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val := val1 + val2 + carry
		lastNode.Val = val % 10
		carry = val / 10

		if l1 != nil || l2 != nil || carry == 1 {
			lastNode.Next = &ListNode{}
			lastNode = lastNode.Next
		}
	}

	return node
}
