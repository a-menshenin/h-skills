package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5, 
						Next: &ListNode{
							Val: 2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	_=head

	head2 := &ListNode{
		Val: 2, 
		Next: &ListNode{
			Val: 1,
			Next: nil,
		},
	}
	_=head2

	head3 := &ListNode{
		Val: 1, 
		Next: nil,
	}

	res := partition(head3, 0)
	_=res
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	currentNode := head
	var (
		LTNodes *ListNode
		currentLTNode *ListNode
		GTNodes *ListNode
		currentGTNode *ListNode
	)

	for {
		if currentNode == nil {
			break
		}
		if currentNode.Val < x {
			if LTNodes == nil {
				LTNodes = currentNode
				currentNode = currentNode.Next

				LTNodes.Next = nil
				currentLTNode = LTNodes

				continue
			}

			currentLTNode.Next = currentNode
			currentLTNode = currentNode

			currentNode = currentNode.Next
			currentLTNode.Next = nil
		} else {
			if GTNodes == nil {
				GTNodes = currentNode
				currentNode = currentNode.Next

				GTNodes.Next = nil
				currentGTNode = GTNodes

				continue
			}

			currentGTNode.Next = currentNode
			currentGTNode = currentNode

			currentNode = currentNode.Next
			currentGTNode.Next = nil
		}
	}

	if currentLTNode != nil {
		currentLTNode.Next = GTNodes
	}

	if LTNodes != nil {
		return LTNodes
	} else if GTNodes != nil {
		return GTNodes
	} else {
		return nil
	}
}
