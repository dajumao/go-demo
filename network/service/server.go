package main

import (
	"fmt"
	"net"
	"php-demo/network/common"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte,1024)
		//fmt.Println("服务器在等待客户端的输入",conn.RemoteAddr().String())
		n,err :=conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出的")
			return
		}
		user := common.JsonDecode(buf[:n])
		fmt.Println(user.Information)
	}
}

func main()  {
	fmt.Println("服务器开始监听了。。。。")
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()
	for  {
		fmt.Println("等待客户端链接")
		conn , err :=listen.Accept()
		if err != nil {
			fmt.Println("accpet(),", err)
			return
		}else {
			fmt.Printf("客户端地址ip: %v\n", conn.RemoteAddr().String())
		}
		go process(conn)
	}
	fmt.Println("listen suc=%v",listen)
}
