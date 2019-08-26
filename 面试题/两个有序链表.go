package main

import "fmt"

/*
两个有序链表合并成一个有序链表
*/

type List struct {
	value int
	next  *List
}

/*
递归实现
*/
func Sort(listO,listT *List) *List {
	if listT == nil && listO ==nil {
		return nil
	}
	if listO == nil {
		return listT
	}
	if listT == nil {
		return listO
	}
	var listTotal *List
	if listO.value < listT.value {
		listTotal =listO
		listTotal.next = Sort(listO.next,listT)
	}else {
		listTotal = listT
		listTotal.next = Sort(listO,listT.next)
	}
	return listTotal
}

/*
非递归实现
*/

func SortFei(listO,listT *List) *List {
	if listT == nil && listO ==nil {
		return nil
	}
	if listO == nil {
		return listT
	}
	if listT == nil {
		return listO
	}
	var res *List
	var sum []int
	for listO != nil && listT != nil  {
		if listO.value < listT.value {
			sum = append(sum,listO.value)
			listO = listO.next
		}else {
			sum = append(sum,listT.value)
			listT = listT.next
		}
	}
	for listO != nil  {
		sum = append(sum,listO.value)
		listO = listO.next
	}
	for listT != nil {
		sum = append(sum,listT.value)
		listT = listT.next
	}
	for i := len(sum)-1; i>=0 ;i--  {
		tem := &List{
			value:sum[i],
			next:res,
		}
		res = tem
	}
	return res
}


func main()  {
	listO := &List{
		value:4,
		next:&List{
			value:6,
			next:&List{
				value:9,
			},
		},
	}

	listT := &List{
		value:1,
		next:&List{
			value:7,
			next:&List{
				value:8,
			},
		},
	}
	total := Sort																																																																																																																																								(listO,listT)
	fmt.Println(total.next.next.value)
}
