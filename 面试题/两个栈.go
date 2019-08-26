package main

import "fmt"

/*
两个栈实现一个队列
*/

/*
实现栈
*/
var strack []int
func push(value int)  {
	strack = append(strack,value)
}

func pull() int  {
	temp := strack[len(strack)-1]
	strack = strack[:len(strack)-1]
	return temp
}

var strackTwo []int
func pushTwo(value int)  {
	strackTwo = append(strackTwo,value)
}

func pullTwo() int  {
	temp := strackTwo[len(strackTwo)-1]
	strackTwo = strackTwo[:len(strackTwo)-1]
	return temp
}

/*
队列的实现
*/
func queuePush(value int)  {
	push(value)
}

func queuePull() int {
	if len(strack) == 0 &&len(strackTwo) == 0 {
		return -1
	}
	if strackTwo ==nil||len(strackTwo) == 0 {
		for len(strack) != 0 {
			temp :=pull()
			pushTwo(temp)
		}
	}
	temp := pullTwo()
	return temp
}


func main()  {
	queuePush(9)
	queuePush(8)
	fmt.Println(queuePull())
	queuePush(7)
	fmt.Println(queuePull())
	fmt.Println(queuePull())
}