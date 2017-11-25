package main

import (
	"github.com/niftynei/algos/three"
	"fmt"
)

func main() {

	//ll := three.ReverseWords("this is a sen tent")
	ll := three.ReverseWords("hello there monkey")
	for ll != nil {
		fmt.Printf("%s\n", ll.Val)
		ll = ll.Next
	}
}

