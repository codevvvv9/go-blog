package main

import (
	"go-blog/config"
	"go-blog/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

// Date 格式化的时间
func Date(layout string) string {
	return time.Now().Format(layout)
}

// ! 小写不对外暴露
func handleIndex(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	path := config.Cfg.System.CurrentDir
	log.Println("path: ", path)
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{
		"date":        Date,
		"isODD":       IsODD,
		"getNextName": GetNextName,
	})
	//templateContent, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析错误：", err)
	}

	//页面上涉及到的所有的数据，必须有定义
	var categories = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	posts := []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张三",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var homeRes = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      posts,
		Total:      1,
		Page:       1,
		Pages:      []int{1},
		PageEnd:    true,
	}
	t.Execute(w, homeRes)
}

func main() {
	//! 程序入口，一个项目只能有一个入口
	// web应用
	server := http.Server{
		Addr: ":8080", //其实就是 127.0.0.1

	}
	// 为了能够访问，需要加上路由监听根路径 就是 localhost:8080
	http.HandleFunc("/", handleIndex)
	// 使用静态服务器处理静态文件
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
