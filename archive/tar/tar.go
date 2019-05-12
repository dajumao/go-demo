package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"os"
)

func main()  {
	fileInfor,err := os.Stat("../archive.md");
	if err != nil {
		fmt.Println("访问的文件不存在")
	}
	fmt.Println("文件修改的时间",fileInfor.ModTime()) 	//文件修改的时间
	fmt.Println(fileInfor.Name())
	fmt.Println(fileInfor.IsDir());		//文件是否一个目录
	file,err := tar.FileInfoHeader(fileInfor,"")
	if err != nil {
		return
	}
	fmt.Println(file.FileInfo().Size())
	//f,err := ioutil.ReadDir("../archive.md")
	//if err != nil {
	//	return
	//}





	// Create a buffer to write our archive to.
	//buf := new(bytes.Buffer)
	//
	//// Create a new tar archive.
	//tw := tar.NewWriter(buf)
	//
	//// Add some files to the archive.
	//var files = []struct {
	//	Name, Body string
	//}{
	//	{"readme.txt", "This archive contains some text files."},
	//	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	//	{"todo.txt", "Get animal handling licence."},
	//}
	//for _, file := range files {
	//	hdr := &tar.Header{
	//		Name: file.Name,
	//		Mode: 0600,
	//		Size: int64(len(file.Body)),
	//	}
	//	if err := tw.WriteHeader(hdr); err != nil {
	//		log.Fatalln(err)
	//	}
	//	if _, err := tw.Write([]byte(file.Body)); err != nil {
	//		log.Fatalln(err)
	//	}
	//}
	//// Make sure to check the error on Close.
	//if err := tw.Close(); err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// Open the tar archive for reading.
	//r := bytes.NewReader(buf.Bytes())
	//tr := tar.NewReader(r)
	//
	//// Iterate through the files in the archive.
	//for {
	//	hdr, err := tr.Next()
	//	if err == io.EOF {
	//		// end of tar archive
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Printf("Contents of %s:\n", hdr.Name)
	//	if _, err := io.Copy(os.Stdout, tr); err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Println()
	//}
}

func Compress(files []*os.File,dest string)  {
	d,_ := os.Create("a.tar.gz")
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _,file := range files {

	}
}
