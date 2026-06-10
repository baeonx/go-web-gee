package main

import (
    "gee"
    "net/http"
    "fmt"
    "io"
    "encoding/json"
    "log"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    r := gee.New()
    
    // 1. GET 测试
    r.GET("/users/:id", func(c *gee.Context) {
        id := c.Param("id")
        log.Printf("GET /users/%s called", id)
        
        c.JSON(http.StatusOK, User{
            ID:   id,
            Name: "John Doe",
            Age:  30,
        })
    })
    
    // 2. POST 测试 - 手动解析
    r.POST("/users", func(c *gee.Context) {
        log.Println("POST /users called")
        
        // 读取请求体
        body, err := io.ReadAll(c.Req.Body)
        if err != nil {
            log.Printf("Error reading body: %v", err)
            c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
            return
        }
        defer c.Req.Body.Close()
        
        log.Printf("Raw body: %s", string(body))
        
        // 解析 JSON
        var user User
        if err := json.Unmarshal(body, &user); err != nil {
            log.Printf("JSON parse error: %v", err)
            c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
            return
        }
        
        log.Printf("Parsed user: Name=%s, Age=%d", user.Name, user.Age)
        
        // 设置 ID
        user.ID = "999"
        
        c.JSON(http.StatusCreated, user)
    })
    
    // 3. PUT 测试
    r.PUT("/users/:id", func(c *gee.Context) {
        id := c.Param("id")
        log.Printf("PUT /users/%s called", id)
        
        body, err := io.ReadAll(c.Req.Body)
        if err != nil {
            log.Printf("Error reading body: %v", err)
            c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
            return
        }
        defer c.Req.Body.Close()
        
        log.Printf("Raw body: %s", string(body))
        
        var user User
        if err := json.Unmarshal(body, &user); err != nil {
            log.Printf("JSON parse error: %v", err)
            c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
            return
        }
        
        log.Printf("Parsed user: Name=%s, Age=%d", user.Name, user.Age)
        
        user.ID = id
        c.JSON(http.StatusOK, user)
    })
    
    // 4. DELETE 测试
    r.DELETE("/users/:id", func(c *gee.Context) {
        id := c.Param("id")
        log.Printf("DELETE /users/%s called", id)
        
        c.JSON(http.StatusOK, map[string]string{
            "message": fmt.Sprintf("User %s deleted successfully", id),
        })
    })
    
    // 5. 添加一个简单的测试路由
    r.GET("/test", func(c *gee.Context) {
        c.String(http.StatusOK, "Test route works!")
    })
    
    log.Println("=== Server starting on :8080 ===")
    log.Println("Available routes:")
    log.Println("  GET    /test")
    log.Println("  GET    /users/:id")
    log.Println("  POST   /users")
    log.Println("  PUT    /users/:id")
    log.Println("  DELETE /users/:id")
    
    r.Run(":8080")
}