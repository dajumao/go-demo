package main

import (
	"io/ioutil"
	"net/http"
	"photoshow-master/simple"
	"regexp"
)

func Show(w http.ResponseWriter,r *http.Request)  {
	f,err := ioutil.ReadFile("./index.html")
	if err != nil {
		return
	}
	w.Write(f)
}

func Choice(w http.ResponseWriter, r * http.Request)  {
	switch  {
	case regexp.MustCompile("/img/*").MatchString(r.URL.Path):
		simple.Picture(w,r)
	default:
		Show(w,r)
	}
}

func main()  {
	http.HandleFunc("/",Choice)
	http.HandleFunc("/simple",simple.SimpleReceive)
	http.ListenAndServe(":8090",nil)
}
