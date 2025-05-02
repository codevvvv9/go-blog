## 初始化主页
main.go程序入口
1. 创建 server，http.server
2. 创建路由监听，http.HandleFunc('/', func(w, r *) {})
3. 启动服务 server.ListenAndServe

go的 template 模板可以直接用，支持双大括号语法{{}}
1. 创建模板目录，引入模板 http/template
2. 组装模板目录，使用 os.Getwd() + '/template/index.html'
3. 解析这个文件生成 templateContent: template.ParseFiles(templatePath)
注意这个 ParseFiles 可以拼接多个文件，很好用
4. 执行templateContent.Execute(writer, data)

## 引入 template 模板和 public资源
/template目录存放所有页面模板
/public目录存放静态文件资源
## 配置目录
/config
config.go数据定义，通过配置文件读取
然后在 init方法中实现一些配置，当系统一启动就会注册配置
## 数据模型目录
/models目录
## 静态文件服务器
```go
http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
```