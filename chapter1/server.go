package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func NewServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

var mu sync.Mutex
var count int

func NewServer2() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count", counter1)
	http.HandleFunc("/from", from1)
	//输出gif
	http.HandleFunc("/getgif", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//解析http request 请求
func from1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//lock 	用于锁住当前请求
func handler1(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)

}

func counter1(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}
