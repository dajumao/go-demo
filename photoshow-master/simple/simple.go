package simple

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func SimpleReceive(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(32<<20)
	if r.MultipartForm == nil {
		return
	}
	file := r.MultipartForm.File["files"][0]
	var paths  = "./img"
	_,err := os.Stat(paths)
	if err != nil {
		err = os.Mkdir("./img",os.ModePerm)
		if err != nil {
			return
		}
	}
	fmt.Println(RandInt(1,12))
	filename := "img"+strconv.Itoa(RandInt(1,12))+".jpg"
	paths = paths+"/"+filename
	f,err := os.OpenFile(paths, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fileUpload,_ := file.Open()
	io.Copy(f,io.Reader(fileUpload))
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func Picture(w http.ResponseWriter,r *http.Request)  {
	r.ParseMultipartForm(32<<20)
	path := string([]byte(r.URL.Path)[1:])
	w.Header().Set("Content-Type","image/png")
	w.Header().Set("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"",path))
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("没有图片")
		return
	}
	w.Write(file)
}

