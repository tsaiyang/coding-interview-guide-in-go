package chapter2

import (
	"coding-interview-guide-in-go/utils"
	"fmt"
)

func PrintCommonPart(h1, h2 *utils.Node) {
	for h1 != nil && h2 != nil {
		if h1.Value < h2.Value {
			h1 = h1.Next
		} else if h1.Value > h2.Value {
			h2 = h2.Next
		} else {
			fmt.Println(h1.Value)
			h1 = h1.Next
			h2 = h2.Next
		}
	}
}
