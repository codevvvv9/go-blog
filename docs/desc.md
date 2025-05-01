## 初始化主页
main.go程序入口
1. 创建 server，http.server
2. 创建路由监听，http.HandleFunc('/', func(w, r *) {})
3. 启动服务 server.ListenAndServe

go的 template 模板可以直接用，支持双大括号语法{{}}
1. 创建模板目录，引入模板 http/template
2. 组装模板目录，使用 os.Getwd() + '/template/index.html'
3. 解析这个文件生成 templateContent: template.ParseFiles(templatePath)
4. 执行templateContent.Execute(writer, data)