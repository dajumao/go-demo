package main

import (
	"fmt"
)

/**
红黑树的定义
 */
type MyRBnode struct {
	color 	bool
	value 	int
	totalNode 	int
	parent  *MyRBnode
	left    *MyRBnode
	right   *MyRBnode
}

var root *MyRBnode
var size int
var RED  bool= true
var BLACK bool=false

func init()  {
	root = nil
	size = 0
}
/*
左旋转
*/
func leftRotate(node *MyRBnode)  {
	x := node.right
	node.right = x.left
	if x.left != nil {
		x.left.parent = node
	}
	x.parent = node.parent
	if node.parent == nil {
		root = x
	}else {
		if node.parent.left == node {
			node.parent.left = x
		}else {
			node.parent.right = x
		}
	}
	x.left = node
	node.parent = x
}

/**
颜色翻转
 */
func flipColors(node *MyRBnode)  {
	node.color = RED
	node.left.color , node.right.color = BLACK,BLACK
}

/*
右旋转
*/
func rightRotate(node *MyRBnode) *MyRBnode {
	x  :=  node.left
	node.right = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

/*
当前节点是否为红结点
*/
func isRed(node *MyRBnode) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

/**
红黑树的添加
 */
func add(value int)  {
	root = addNode(root,value)
	root.color = BLACK
}

/*
向node为根元素的红黑树插入元素
*/
func addNode(node *MyRBnode,value int) *MyRBnode {
	if node == nil {
		size++
		node = &MyRBnode{
			value:value,
			color:RED,
		}
		return node
	}
	if value < node.value {
		node.left = addNode(node.left,value);
	}else if value >node.value {
		node.right = addNode(node.right,value)
	}else {
		node.value = value
	}

	/*
	左旋转判断
	*/
	//if isRed(node.right) && !isRed(node.left) {
	//	node = leftRotate(node)
	//}
	/*
	右旋转判断
	*/
	//if isRed(node.left) && isRed(node.left.left) {
	//	node = rightRotate(node)
	//}
	/*
	判断是否需要颜色翻转
	*/
	//if isRed(node.left) && isRed(node.right) {
	//	flipColors(node)
	//}

	return node
}


func main()  {
	//redTree := false
	//blackTree:=true
	add(2)
	add(4)
	add(1)
	add(3)
	fmt.Println(root.color)
}
