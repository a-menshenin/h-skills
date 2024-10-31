package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(l1, l2)

	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	carry := false
	tmp := res
	for {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2
		if carry {
			sum++
			carry = false
		}

		if v := sum % 10; sum >= 10 && v >= 0 {
			sum = v
			carry = true
		}

		tmp.Val = sum

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		tmp.Next = &ListNode{}

		if l1 == nil && l2 == nil {
			if carry {
				tmp.Next.Val = 1
				tmp.Next.Next = nil
			} else {
				tmp.Next = nil
			}

			break
		}

		tmp = tmp.Next
	}

	return res
}
