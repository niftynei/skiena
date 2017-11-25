package three

import (
	"fmt"
)

type Node struct {
	Left *Node
	Right *Node
	Val int
}

type LNode struct {
	Next *LNode
	Val int
}

func (l *LNode) Add(val int) *LNode {
	l.Next = &LNode{Val: val}
	return l.Next
}

func (n *Node) insert(val int) {
	if n.Val == 0 {
		n.Val = val
		return;
	}
	if n.Val == val {
		return;
	}
	if val < n.Val {
		if n.Left == nil {
			n.Left = &Node{}
		}
		n.Left.insert(val)
	} else {
		if n.Right == nil {
			n.Right = &Node{}
		}
		n.Right.insert(val)
	}
}

func (n *Node) Insert(node *Node) {
	if n.Val > node.Val {
		if n.Left == nil {
			n.Left = node
			return
		} else {
			n.Left.Insert(node)
		}
	} else {
		if n.Right == nil {
			n.Right = node
			return
		} else {
			n.Right.Insert(node)
		}
	}
}

func (n *Node) Min() *Node {
	node := n
	for node.Left != nil {
		node = n.Left
	}
	return node
}

// have to 'rebalance' after deletion
func Del(n *Node, val int) bool {
	if n.Val == val {
		// delete & rebal
		if n.Left != nil {
			tmp := n.Left
			if tmp.Right != nil {
				if n.Right != nil {
					if n.Right.Val > tmp.Right.Val {
						n.Right.Insert(tmp.Right)
						tmp.Right = n.Right
					}
				}
			} else if n.Right != nil {
				tmp.Right = n.Right
			}
			n = tmp
		} else if n.Right != nil {
			n = n.Right
		} else {
			n = nil
		}
		return true
	}

	if val < n.Val {
		if n.Left != nil {
			return Del(n.Left, val)
		} else {
			return false
		}
	} else {
		if n.Right != nil {
			return Del(n.Right, val)
		} else {
			return false
		}
	}
}

func (tree *Node) ToList() *LNode {
	start, _ := list(tree, nil, nil, false)
	return start
}

func (tree *Node) ToOrderedList() *LNode {
	start, _ := list(tree, nil, nil, true)
	return start
}

func list(node *Node, link *LNode, start *LNode, order bool) (*LNode, *LNode) {
	nextLink := link

	if !order {
		if nextLink == nil {
			nextLink = &LNode{Val:node.Val}
			start = nextLink
		} else {
			nextLink = nextLink.Add(node.Val)
		}
	}

	if node.Left != nil {
		start, nextLink = list(node.Left, nextLink, start, order)
	}

	if order {
		if nextLink == nil {
			nextLink = &LNode{Val:node.Val}
			start = nextLink
		} else {
			nextLink = nextLink.Add(node.Val)
		}
	}

	if node.Right != nil {
		start, nextLink = list(node.Right, nextLink, start, order)
	}


	return start, nextLink
}

func Same(t1 *Node, t2 *Node) bool {
	arr := []int{}
	arr2 := []int{}
	TreeToArr(t1, &arr)
	TreeToArr(t2, &arr2)

	if len(arr) != len(arr2) {
		return false
	}
	for i, val := range arr {
		if val != arr2[i] {
			return false
		}
	}

	return true
}

func TreeToArr(n *Node, arr *[]int) {
	if n.Val != 0 {
		*arr = append(*arr, n.Val)
	}
	if n.Left != nil {
		TreeToArr(n.Left, arr)
	}
	if n.Right != nil {
		TreeToArr(n.Right, arr)
	}
}

func (n *Node) PrintTree() {
	if n.Left != nil {
		n.Left.PrintTree()
	}
	if n.Val != 0 {
		fmt.Printf("val:%d\n", n.Val)
	}
	if n.Right != nil {
		n.Right.PrintTree()
	}
}

func MakeTree(arr []int) (n *Node) {
	node := &Node{}
	for _, val := range arr {
		node.insert(val)
	}

	return node
}
