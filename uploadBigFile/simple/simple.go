package simple

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SimpleReceive(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(32<<20)
	if r.MultipartForm == nil {
		return
	}
	file := r.MultipartForm.File["files"][0]
	var paths  = "./newFile"
	_,err := os.Stat(paths)
	if err != nil {
		err = os.Mkdir("./newFile",os.ModePerm)
		if err != nil {
			return
		}
	}
	f,err := os.OpenFile(paths+"/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fileUpload,_ := file.Open()
	io.Copy(f,io.Reader(fileUpload))
	fmt.Println(file.Filename)
	fmt.Println("----------------")
	fmt.Fprintln(w,"upload ok!")
}