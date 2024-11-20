package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	second := &ListNode{
		Val: 2,
		Next: nil,
	}

	second.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: -4,
			Next: second,
		},
	}

	head := &ListNode{
		Val: 3,
		Next: second,
	}
	_=head

	test2 := &ListNode{
		Val: 1,
		Next: nil,
	}
	_=test2

	res := hasCycle(head)
	fmt.Println(res)
}

// https://leetcode.com/problems/linked-list-cycle/submissions/1458179396/
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	if head == nil || head.Next == nil {
		return false
	}

	for slow != nil && fast != nil{
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next

		if fast == slow {
			return true
		}
	}

	return false
}
