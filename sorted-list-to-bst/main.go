package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	listNodes := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
						Next: nil,
					},
				},
			},
		},
	}
	_=listNodes
	listNodes2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res := sortedListToBST(listNodes2)
	_=res

	fmt.Println(res)
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// Алгоритм:
	// 1) Возьмите середину связанного списка и сделайте ее корневой.
	// 2) Рекурсивно проделайте то же самое для левой и правой половины.
	listLen := 0
	tmpList := head
	// Определяем длину списка
	for tmpList != nil {
		listLen++
		if tmpList.Next == nil {
			break
		}
		tmpList = tmpList.Next
	}


	var (
		leftSubtree *ListNode
		rightSubtree *ListNode
		curLeftSubtree *ListNode
	)
	tmpList = head
	for i := 0; i < listLen / 2; i++ {
		if leftSubtree == nil {
			leftSubtree = &ListNode{
				Val: tmpList.Val,
			}
			curLeftSubtree = leftSubtree
		} else {
			curLeftSubtree.Next = tmpList
			curLeftSubtree = curLeftSubtree.Next
		}
		tmpList = tmpList.Next
	}

	midNode := tmpList

	if curLeftSubtree != nil {
		curLeftSubtree.Next = nil
	}
	rightSubtree = tmpList.Next
	midNode.Next = nil

	// Делаем середину связанного списка корневой
	res := &TreeNode{
		Val: midNode.Val,
	}
		
	// Рекурсивно проделайте то же самое для левой и правой половины
	leftNode := sortedListToBST(leftSubtree)
	rightNode := sortedListToBST(rightSubtree)

	res.Left = leftNode
	res.Right = rightNode

    return res
}
