package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}
/*
后序遍历
*/
func diguiTree(node *node)  {
	if node.left != nil {
		diguiTree(node.left)
	}
	if node.right != nil {
		diguiTree(node.right)
	}
	fmt.Println(node.value)
}

/*
后序遍历的非递归
*/

func feidiguiTree(nodes *node)  {
	var stack []*node
	var flag *node
	for nodes != nil {
		stack = append(stack,nodes)
		nodes = nodes.left
	}
	for len(stack) != 0 {
		nodes = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nodes.right == nil||nodes.right == flag {
			fmt.Println(nodes.value)
			flag = nodes
		}else {
			stack = append(stack,nodes)
			nodes = nodes.right
			for nodes!=nil {
				stack = append(stack,nodes)
				nodes = nodes.left
			}
		}
	}
}

func main()  {
	tree := &node{
		value:4,
		left:&node{
			value:9,
			right:&node{
				value:5,
			},
		},
		right:&node{
			value:4,
			left:&node{
				value:1,
			},
		},
	}
	diguiTree(tree)
	fmt.Println("=========================")
	feidiguiTree(tree)
}

func j(h *node)  {
	var stack []*node
	var flag *node
	for h != nil {
		stack = append(stack,h)
		h = h.left
	}
	for len(stack)>0 {
		d := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if h.right == nil|| h.right == flag {
			fmt.Println(h.value)
			flag = d
		}else {
			stack = append(stack,h)
			h = h.right
			for h != nil {
				stack = append(stack,h)
				h = h.left
			}
		}
	}
}