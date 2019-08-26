package main

import "fmt"

var pushstack []int
var popstack []int

func push(number int)  {
	pushstack = append(pushstack,number)
}

func pop() int  {
	if (popstack == nil && pushstack == nil) ||(len(popstack) == 0 && len(pushstack) == 0) {
		panic("队列为空")
	}
	if len(pushstack) != 0 && len(popstack) == 0 {
		for i:=len(pushstack)-1;i>=0;i-- {
			popstack = append(popstack,pushstack[i])
		}
		pushstack = pushstack[:0]
	}
	value := popstack[len(popstack)-1]
	popstack = popstack[:len(popstack)-1]
	return value
}

func main()  {
	push(0)
	push(9)
	fmt.Println(pop())
	fmt.Println(pop())
	push(8)
	fmt.Println(pop())
}
