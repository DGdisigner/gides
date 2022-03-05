package server

import (
	"fmt"
	"gides/src/driver"
	"log"
	"net/http"
)

type Server struct {
	driver.CacheDriver
	driver.Cluster
	Ip   string
	Port string
}

func (s *Server) run() {
	http.HandleFunc("/get", s.get)
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", s.Ip, s.Port), nil)
	if err != nil {
		log.Println(err.Error())
	}
}

// handler函数
func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	_, err := w.Write([]byte("www.5lmh.com"))
	if err != nil {
		log.Println(err.Error())
	}
}
