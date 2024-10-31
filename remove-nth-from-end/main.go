package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: & ListNode{
						Val: 5, 
						Next: nil,
					},
				},
			},
		},
	}
	_=head1

	head2 := &ListNode{
		Val: 1,
		Next: nil,
	}
	_=head2

	head3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}
	_=head3

	res := removeNthFromEnd(head1, 2)
	_=res
	fmt.Println("")
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	tmp := head
	listLen := 1
	for tmp.Next != nil {
		tmp = tmp.Next
		listLen++
	}

	nextCount := listLen - n - 1
	if nextCount < 0 {
		head = head.Next
	} else {
		first := head
		second := head.Next
		for i := 0; i < nextCount; i++ {
			if second == nil {
				break
			}
			first = second
			second = second.Next
		}

		if second != nil {
			first.Next = second.Next
		} else {
			head = nil
		}
	}

	return head
}
