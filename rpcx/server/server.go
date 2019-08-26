package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/*
简易
*/
//type Cqh struct {
//
//}
//
//func (p *Cqh)Test(request string,reply *string)error  {
//	*reply="test"+request
//	return nil
//}
//
//func main()  {
//	rpc.RegisterName("Cqh",new(Cqh))
//	listener,err := net.Listen("tcp",":1234")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		conn,err := listener.Accept()
//		if err != nil {
//			log.Fatal("listener accept",err)
//		}
//		rpc.ServeConn(conn)
//	}
//}

type Panda int

func (this *Panda)Getinfo(argType int,replyType *int)error  {
	*replyType=1+argType
	return nil
}

func pandatext(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"panda")
}

func main()  {
	http.HandleFunc("/",pandatext)
	pd := new(Panda)
	rpc.Register(pd)
	rpc.HandleHTTP()
	ln,err := net.Listen("tcp",":10086")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(ln,nil)
}
