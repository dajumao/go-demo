package main

import "fmt"

type student struct{
	Name string
}

func main()  {
	//gen:= func(ctx context.Context) <-chan int{
	//	dst:=make(chan int)
	//	n:=1
	//	go func() {
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				return
	//			case dst<-n:
	//				n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//ctx,cancel := context.WithCancel(context.Background())
	//defer cancel()
	//for n:=range gen(ctx) {
	//	fmt.Println(n)
	//	if n==5 {
	//		break
	//	}
	//}

	//type favContextKey string
	//
	//f := func(ctx context.Context, k favContextKey) {
	//	if v := ctx.Value(k); v != nil {
	//		fmt.Println("found value:", v)
	//		return
	//	}
	//	fmt.Println("key not found:", k)
	//}
	//
	//k := favContextKey("language")
	//ctx := context.WithValue(context.Background(), k, "Go")
	//
	//f(ctx, k)
	//f(ctx, favContextKey("color"))
	m := make(map[string]*student)
	stus := []student{
		{Name:"z"},
		{Name:"r"},
	}
	for _,stu:=range stus{
		j:=stu
		m[stu.Name] = &j
	}
	for _,v := range m  {
		fmt.Println(v)
	}
}
