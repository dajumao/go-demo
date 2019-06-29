package main

import (
	"fmt"
	"net"
)

func main() {
	//建立连接
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	defer conn.Close()
	if err != nil{
		fmt.Println(err)
		return
	}

	go func(){
		var strbuf string
		for{
			fmt.Print("输入要发送的数据: ")
			fmt.Scanf("%s",&strbuf)
			if err != nil{
				fmt.Println(err)
				break
			}

			//write
			_,err = conn.Write([]byte(strbuf))	//\r\n
			if err != nil{
				fmt.Println(err)
				break
			}
		}
	}()


	//read
	buffer := make([]byte,1024)
	for{
		n,err := conn.Read(buffer)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println("收到服务器的响应数据: ",string(buffer[:n]))
	}
}