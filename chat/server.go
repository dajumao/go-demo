package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	lisener,err := net.Listen("tcp","127.0.0.1:8888")
	defer lisener.Close()
	if err != nil{
		fmt.Println(err)
		return
	}

	for{
		//阻塞接受连接
		conn,err := lisener.Accept()	//blocked
		if err != nil{
			fmt.Println(err)
			conn.Close()
			continue
		}

		//协程处理连接
		go func(conn net.Conn){
			defer conn.Close()

			cli_addr := conn.RemoteAddr().String()
			fmt.Println("cli_addr: ",cli_addr)

			for{
				//read
				buf := make([]byte,1024)
				n,err := conn.Read(buf)
				if err != nil{
					fmt.Println(err)
					return
				}
				fmt.Println("收到客户端的请求数据: ",string(buf[:n]), "长度: ", n)

				if string(buf[:n]) == "exit"{
					fmt.Println("收到客户端断开连接的请求，即将断开连接")
					return
				}

				//write
				buf = []byte(strings.ToUpper(string(buf[:n])))
				_,err = conn.Write(buf[:n])
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}(conn)
	}
}

