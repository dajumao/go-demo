package main

import "fmt"

/*
两个队列实现一个栈
 */
var queueO []int
func queuePushO(value int)  {
	queueO = append(queueO,value)
}

func queuePullO() int  {
	if queueO == nil || len(queueO) == 0 {
		return -1
	}
	temp := queueO[0]
	queueO = queueO[1:]
	return temp
}

var queueT []int
func queuePushT(value int)  {
	queueT = append(queueT,value)
}

func queuePullT() int  {
	if queueT == nil || len(queueT) == 0 {
		return -1
	}
	temp := queueT[0]
	queueT = queueT[1:]
	return temp
}

/*
栈实现
*/
func pushS(value int)  {
	if len(queueO) == 0 && len(queueT) == 0 {
		queuePushO(value)
		return
	}
	if len(queueO) != 0 {
		queuePushO(value)
		return
	}
	if len(queueT) != 0 {
		queuePushT(value)
		return
	}
}

func pullS() int {
	if len(queueO) == 0 && len(queueT) == 0 {
		return -1
	}
	var tmp int
	if len(queueO) == 0 {
		for len(queueT)>1 {
			queuePushO(queuePullT())
		}
		tmp = queuePullT()
		return tmp
	}
	if len(queueT) == 0 {
		for len(queueO)>1 {
			queuePushT(queuePullO())
		}
		tmp = queuePullO()
		return tmp
	}
	return tmp
}

func main()  {
	pushS(4)
	pushS(7)
	pushS(9)
	fmt.Println(pullS())
	pushS(0)
	fmt.Println(pullS())
}