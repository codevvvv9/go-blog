## 初始化主页
main.go程序入口
1. 创建 server，http.server
2. 创建路由监听，http.HandleFunc('/', func(w, r *) {})
3. 启动服务 server.ListenAndServe