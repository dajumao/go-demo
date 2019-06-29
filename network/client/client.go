package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"php-demo/network/common"
	"strings"
	"time"
)

func main()  {

	conn,err := net.Dial("tcp",common.GetIp()+":8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	fmt.Println("conn 成功")
	reader := bufio.NewReader(os.Stdin)
	jsonIn := common.ChatInformation{}
	for {
		line ,err :=reader.ReadString('\n')
		if err != nil{
			fmt.Println("reading=",err)
			break
		}
		line =strings.Trim(line,"\r\n")
		if  line == "exit" {
			fmt.Println("客户端退出了")

		}
		jsonIn.Information = line
		jsonIn.Time		   = time.Now()
		_, err = conn.Write(common.JsonCode(&jsonIn))
		if err != nil {
			fmt.Println("conn.write",err)
			break
		}
	}
	fmt.Println("客户端退出")
}
