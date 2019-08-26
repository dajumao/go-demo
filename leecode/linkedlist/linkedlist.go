package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

func list(nodee *Node) *Node {
	if nodee == nil {
		return nodee
	}
	var temp *Node = nil
	for nodee.next != nil {
		g := &Node{
			value:nodee.value,
		}
		g = temp
		temp = g
		nodee = nodee.next

	}
	return temp
}

func main()  {
	v := &Node{
		3,
		&Node{
			4,
			&Node{
				5,
				nil,
			},
		},
	}
	v = list(v)
	for v.next!=nil {
		fmt.Println(v.value)
		v = v.next
	}
	fmt.Println(v.value)
}
