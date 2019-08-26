package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"php-demo/uploadBigFile/simple"
	"php-demo/uploadBigFile/additionalUpload"
)

/**
io读取模板
*/
func Show(w http.ResponseWriter,r *http.Request)  {
	f,err := ioutil.ReadFile("./html/file.html")
	if err != nil {
		return
	}
	w.Write(f)
}
/**
template模板
 */
func ShowTemplate(w http.ResponseWriter,r *http.Request)  {
	t,err := template.ParseFiles("./html/file.html")
	if err != nil {
		log.Println("无法访问模板")
	}
	t.Execute(w,"")
}

/**
防止当其中一个路由发生错误ｐａｎｉｃ时,整个程序崩溃
 */
func loginPanic(handle http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[%v] caught panic:%v",request.RemoteAddr,err)
			}
		}()
		handle(writer,request)
	}
}

func main()  {
	http.HandleFunc("/",loginPanic(Show))
	http.HandleFunc("/sss",loginPanic(ShowTemplate))
	http.HandleFunc("/simple",loginPanic(simple.SimpleReceive))
	http.HandleFunc("/additionalUpload",(additionalUpload.AdditionalUploadReceive))
	http.ListenAndServe(":8088",nil)
}

