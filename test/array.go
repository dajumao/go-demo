package main

import "fmt"

func Find(target int,testmap [][]int)  {
	if testmap == nil || cap(testmap) == 0 {
		return
	}
	for i := 0;i<cap(testmap) ; i++{
		fmt.Println(testmap[i][cap(testmap)-1])
		if testmap[i][cap(testmap)-1] >= target {
			low := 0
			hight := cap(testmap[0])
			var mid int
			for low <=hight  {
				mid = (low+hight)/2
				if testmap[i][mid] == target {
					fmt.Println("-------")
					break
				}else if testmap[i][mid] > target {
					hight = mid-1
				}else {
					low = mid+1
				}
			}
		}
	}
}

func main()  {
	target := 4
	testmap:= [][]int{
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}
	Find(target,testmap)
}
