package three

type WNode struct {
	Next *WNode
	Val string
}

func ReverseWords(phrase string) *WNode {
	var acc string
	var node *WNode
	for _, c := range phrase {
		if c == ' ' {
			// emit word
			tmp := node
			node = &WNode{}
			node.Next = tmp
			node.Val = acc
			acc = ""
		} else {
			acc += string(c)
		}
	}
	if acc != "" {
		tmp := node
		node = &WNode{}
		node.Val = acc
		node.Next = tmp
	}

	return node
}
