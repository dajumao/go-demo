package common

import (
	"fmt"
	"net"
	"strings"
)

func GetIp() string {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
