package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	
	// GET 接口
	r.GET("/user", func(c *gee.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":   1,
			"name": "张三",
		})
	})
	
	// POST 接口
	r.POST("/user", func(c *gee.Context) {
		c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "创建成功",
		})
	})
	
	// PUT 接口
	r.PUT("/user/1", func(c *gee.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "更新成功",
		})
	})
	
	// DELETE 接口
	r.DELETE("/user/1", func(c *gee.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "删除成功",
		})
	})
	
	// 打印接口信息
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
	
	// 启动服务器
	r.Run(":8080")
}