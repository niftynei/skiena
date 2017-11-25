package main

import (
	//"github.com/niftynei/algos/four"
	"fmt"
	"math/rand"
)

func main() {

	for _, value := range rand.Perm(3) {
		fmt.Println(value)
	}
}

