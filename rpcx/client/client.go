package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
简易
*/
//func main()  {
//	client,err := rpc.Dial("tcp",":1234")
//	if err != nil {
//		log.Fatal("dialing",err)
//	}
//	var reply string
//	err = client.Call("Cqh.Test","hello",&reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(reply)
//
//}
func main()  {
	cli,err := rpc.DialHTTP("tcp",":10086")
	if err!= nil {
		log.Fatal(err)
	}
	var val int
	//err = cli.Call("Panda.Getinfo",123,&val)

	divCall := cli.Go("Panda.Getinfo",123,&val,nil)
	//select {
	//case replyCall := <-divCall.Done:
	//	if replyCall.Error !=nil {
	//		log.Fatal(replyCall.Error)
	//	}
	//	fmt.Println(val)
	//
	//}
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Fatal(replyCall.Error)
	}
	fmt.Println(val)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(val)
}