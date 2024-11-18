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

func hasCycle(head *ListNode) bool {
	slowPointer := head
	fastPointer := head
	if head == nil || fastPointer.Next == nil && slowPointer.Next == nil {
		return false
	}

	for {
		if slowPointer == nil || fastPointer == nil {
			return false
		}

		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
		if fastPointer == nil {
			return false
		}
		fastPointer = fastPointer.Next

		if fastPointer == slowPointer {
			return true
		}
	}
}
