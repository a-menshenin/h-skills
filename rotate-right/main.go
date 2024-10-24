package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}

	_ = rotateRight(head, 4)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

    last := head
	listLen := 1
    for last.Next != nil {
        last = last.Next
		listLen++
    }

	if k > listLen {
		k = k % listLen
	}

	prevLast := head
	for prevLast.Next != last {
		prevLast = prevLast.Next
	}

	for i := 0; i < k; i++ {
		last.Next = head
		prevLast.Next = nil

		head = last
		last = prevLast

		prevLast = head
		for prevLast.Next != last {
			prevLast = prevLast.Next
		}
	}

	return head
}
