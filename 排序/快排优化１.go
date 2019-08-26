package main

import (
	"fmt"
)

func division(list[] int,left int,right int) int {
	base:=list[left]
	for left < right  {
		for left<right&&list[right] >base  {
			right--
		}
		list[left] = list[right]
		for left<right&&list[left] < base  {
			left++
		}
		list[right] = list[left]
	}
	list[left] = base
	return left
}

func k(list[] int,left,right int)  {
	if left<right {
		base := division(list,left,right)
		k(list,left,base-1)
		k(list,base+1,right)
	}
}

func QuitSrt(list[] int,left,right int)  {
	if left < right {
		base := division(list,left,right)
		QuitSrt(list,left,base-1)
		QuitSrt(list,base+1,right)
	}
}



/*
快排非递归实现
*/

func QuitSort(list[] int,left,right int)  {
	if left >= right {
		return
	}
	var stack[] int
	stack = append(stack,left,right)
	for len(stack)>0 {
		rightr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftr := stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if leftr < rightr {
			base := division(list,leftr,rightr)
			stack = append(stack,leftr,base-1)
			stack = append(stack,base+1,rightr)
		}
	}
}


/**
三数取中
 */


func main()  {
	arr := []int{-9,6, 4,-8, 8, 9, 2,0, 3, 1}
	QuitSort(arr,0, len(arr)-1);
	for _, value := range arr {
		fmt.Println(value)
	}
}
