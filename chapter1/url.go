package chapter1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)
//获取url data
//Users/lijun/Desktop/go_build_main_go http://gopl.io
func GetDataByUrl1(){
	for _,url:=range  os.Args[1:]{
		resp,err:=http.Get(url)
		if err!=nil {
			fmt.Println("get data error")
			os.Exit(1)
		}
		data,errgetBody:=ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if errgetBody!=nil {
			fmt.Println("get body data error")
			os.Exit(1)
		}
		fmt.Println(data)
	}
}
//并发获取url data
// go_build_main_go http://gopl.io https://godoc.org https://www.douban.com/
func GetDataByuRL2() {
	start:=time.Now()
	ch := make(chan string)
	for _,url:=range  os.Args[1:]{
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string,ch chan<-string ) {
	start:=time.Now()
	resp,err:=http.Get(url)
	if err!=nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes,err:=io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if err!=nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs:=time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)

}