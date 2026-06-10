package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	// GET 接口 - 获取用户信息
	r.GET("/user", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"name":"张三","age":25}`))
	})

	// POST 接口 - 创建用户
	r.POST("/user", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"用户创建成功"}`))
	})

	// PUT 接口 - 更新用户
	r.PUT("/user/1", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"用户更新成功"}`))
	})

	// DELETE 接口 - 删除用户
	r.DELETE("/user/1", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"用户删除成功"}`))
	})

	// 打印启动信息
	println("========================================")
	println("Gee Framework Server Starting...")
	println("Server address: http://localhost:8080")
	println("========================================")
	println("Available endpoints:")
	println("  GET    /user       - 获取用户信息")
	println("  POST   /user       - 创建用户")
	println("  PUT    /user/1     - 更新用户")
	println("  DELETE /user/1     - 删除用户")
	println("========================================")
	println("Test commands:")
	println("  curl http://localhost:8080/user")
	println("  curl -X POST http://localhost:8080/user")
	println("  curl -X PUT http://localhost:8080/user/1")
	println("  curl -X DELETE http://localhost:8080/user/1")
	println("========================================")

	r.Run(":8080")
}