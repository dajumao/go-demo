package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func reConstructBinaryTree(pre[] int,in[] int)  {
	if pre == nil || in == nil {
		return
	}
	tree := &Node{}
	result := constructBinary(pre,in,tree)
	fmt.Println(result.right.left.value)
}
func constructBinary(pre[] int,in[] int,tree *Node) *Node {
	if len(pre) == 0 || len(in) == 0 || tree == nil {
		return tree
	}
	tree.value = pre[0]
	if len(pre) == 1 {
		return tree
	}
	var tmp int
	for i := 0 ; i < len(in) ; i++ {
		if in[i] == pre[0] {
			tmp = i
		}
	}
	if tmp == len(pre)-1 {
		tree.left = &Node{}
		tree.left = constructBinary(pre[1:],in[0:tmp],tree.left)
	} else if tmp == 0 {
		tree.right = &Node{}
		tree.right = constructBinary(pre[1:],in[1:],tree.right)
	} else {
		tree.left = &Node{}
		tree.right = &Node{}
		tree.left = constructBinary(pre[1:tmp+1],in[:tmp],tree.left)
		tree.right = constructBinary(pre[tmp+1:],in[tmp+1:],tree.right)
	}
	return tree
}


func main()  {
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,8,6}
	reConstructBinaryTree(pre,in)
}