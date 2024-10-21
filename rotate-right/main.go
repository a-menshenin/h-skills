package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := head
	for last.Next != nil {
		last = last.Next
	}

	for i := 0; i < k; i++ {
		last.Next = head
		prevLast := head
		for prevLast.Next.Val != last.Val {
			prevLast = prevLast.Next
		}
		prevLast.Next = nil

		head = last
		last = prevLast
	}

	return head
}
