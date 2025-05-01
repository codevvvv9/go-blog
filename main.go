package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// ! 小写不对外暴露
func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "吴少的 go 博客"
	indexData.Desc = "Welcome to the Welcome page."
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}
func main() {
	//! 程序入口，一个项目只能有一个入口
	// web应用
	server := http.Server{
		Addr: ":8080", //其实就是 127.0.0.1

	}
	// 为了能够访问，需要加上路由监听根路径 就是 localhost:8080
	http.HandleFunc("/", handleIndex)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
