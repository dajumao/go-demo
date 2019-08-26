package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Replaces(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s
	}
	if m := strings.Count(s, old); m == 0 {
		return s
	} else if n < 0 || m < n {
		n = m
	}
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return string(t[0:w])
}

func main()  {
	fmt.Println(Replaces("dsadasd dsad sddsad"," ", "%20",-1))
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println(copy(slice2, slice1))
	fmt.Println(slice2)
}
