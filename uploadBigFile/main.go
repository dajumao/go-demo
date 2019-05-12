package main

import (
	"io/ioutil"
	"net/http"
	"php-demo/uploadBigFile/simple"
	"php-demo/uploadBigFile/additionalUpload"
)

func Show(w http.ResponseWriter,r *http.Request)  {
	f,err := ioutil.ReadFile("./html/file.html")
	if err != nil {
		return
	}
	w.Write(f)
}

func main()  {
	http.HandleFunc("/",Show)
	http.HandleFunc("/simple",simple.SimpleReceive)
	http.HandleFunc("/additionalUpload",additionalUpload.AdditionalUploadReceive)
	http.ListenAndServe(":8088",nil)
}
