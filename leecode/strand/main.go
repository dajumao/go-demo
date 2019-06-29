package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	length := 0
	largestLength := 0
	var g []string
	a := strings.Split(s,"")
	for key,value := range a {
		if key == 0 {
			g = append(g,value)
			length++
			largestLength++
			continue
		}
		for keyg,valueg := range g {
			if value == valueg {
				g = g[keyg+1:]
			}
		}
		g = append(g,value)
		length = len(g)
		if length>largestLength {
			largestLength = length
		}
	}
	return largestLength
}

func main()  {
	f := "99908067789456"
	fmt.Println(lengthOfLongestSubstring(f))
}
