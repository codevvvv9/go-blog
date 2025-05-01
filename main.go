package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// ! 小写不对外暴露
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "吴少的 go 博客"
	indexData.Desc = "Welcome to the Welcome page."
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}
func handleIndexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "吴少的 go 博客"
	indexData.Desc = "Welcome to the Welcome page."
	//! 获取项目根据经
	path, _ := os.Getwd()
	templatePath := path + "/template/index.html"

	templateContent, _ := template.ParseFiles(templatePath)
	templateContent.Execute(w, indexData)

}
func main() {
	//! 程序入口，一个项目只能有一个入口
	// web应用
	server := http.Server{
		Addr: ":8080", //其实就是 127.0.0.1

	}
	// 为了能够访问，需要加上路由监听根路径 就是 localhost:8080
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/index.html", handleIndexHtml)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
