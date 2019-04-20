package main

import (
	"bufio"
	"os"
	"php-demo/Reptile/snake"
)

/**
从控制台霍去病输入的csdn首页的地址输入1停止输入
 */
func main()  {
	var url string
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		url = reader.Text()
	}
	snake.Gethttp(url)
}
