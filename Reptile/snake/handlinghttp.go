package snake

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Gethttp(url string)  {
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("你输入的地址有问题")
		return
	}
	defer resp.Body.Close()
	resp.Header.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
	headers := resp.Header
	for  k,v := range headers {
		fmt.Println(k,":",v)
	}
	body,err := ioutil.ReadAll(resp.Body)
	for err != nil  {
		return
	}
	analysisUrl(string(body))
}

func analysisUrl(body string)  {
	
}