package main

import (
	"fmt"
	"log"
	"sync"
)

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表的实现
func reversrList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil  {
		temp := cur.Val
		cur = cur.Next
		f := &ListNode{
			temp,
			pre,
		}
		pre = f
	}
	return pre
}

func printu(list *ListNode)  {
	for  {
		fmt.Println(list.Val)
		if list.Next == nil{
			break
		}
		list = list.Next
	}
}

//func main() {
//	head := new(ListNode)
//	head.Val = 1
//	ln2 := new(ListNode)
//	ln2.Val = 2
//	ln3 := new(ListNode)
//	ln3.Val = 3
//	ln4 := new(ListNode)
//	ln4.Val = 4
//	ln5 := new(ListNode)
//	ln5.Val = 5
//	head.Next = ln2
//	ln2.Next = ln3
//	ln3.Next = ln4
//	ln4.Next = ln5
//
//	list := reversrList(head)
//	printu(list)
//}

func main() {
	var h sync.WaitGroup
	h.Add(1)
	go func() {
		h.Done()
		log.Print("Hello")
	}()
	h.Wait()
	log.Print("Wold")
}