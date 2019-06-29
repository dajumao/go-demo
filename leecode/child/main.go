package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	str := ""
	if nil == strs||len(strs) == 0||len(strs[0]) == 0 {
		return str
	}
Exit:
	for key,value := range strs[0] {
		tmp := value
		for _,values := range strs {
			if len(values) < key+1 {
				return str
			}
			if tmp != int32(values[key]) {
				break Exit
			}
		}
		str = str + string(tmp)
	}
	return str
}

func main()  {
	s :=[]string{
		"aa","a",
	}
	fmt.Println(longestCommonPrefix(s))
}
