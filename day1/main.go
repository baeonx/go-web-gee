package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New() // 创建一个gee实例

	// 首页 - 返回欢迎信息
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome to Gee Framework!"))
	})

	// Hello接口 - 返回问候语
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// 获取用户信息 - 返回JSON格式数据
	r.GET("/user", func(w http.ResponseWriter, req *http.Request) {
		// 设置响应头为JSON格式
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"name":"张三","age":25}`))
	})

	// 创建用户 - POST请求示例
	r.POST("/user", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201 Created
		w.Write([]byte(`{"message":"用户创建成功"}`))
	})

	// 获取用户ID - 带查询参数的接口
	r.GET("/user/profile", func(w http.ResponseWriter, req *http.Request) {
		// 获取URL查询参数 ?name=张三
		name := req.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("User profile"))
		} else {
			w.Write([]byte(fmt.Sprintf("User profile: %s", name)))
		}
	})

	// 健康检查接口 - 常用于服务监控
	r.GET("/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"gee"}`))
	})

	// 打印启动信息
	fmt.Println("=== Gee Framework Test Server ===")
	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("")
	fmt.Println("Available routes:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /hello")
	fmt.Println("  GET  /user")
	fmt.Println("  POST /user")
	fmt.Println("  GET  /user/profile?name=xxx")
	fmt.Println("  GET  /health")
	fmt.Println("")
	fmt.Println("Test with curl:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/hello")
	fmt.Println("  curl http://localhost:8080/user")
	fmt.Println("  curl -X POST http://localhost:8080/user")
	fmt.Println("  curl 'http://localhost:8080/user/profile?name=张三'")
	fmt.Println("  curl http://localhost:8080/health")
	fmt.Println("================================")

	r.Run(":8080")
}