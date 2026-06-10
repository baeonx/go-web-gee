package main

import (
    "gee"
    "net/http"
    "fmt"
)

func main() {
    r := gee.New()
    
    // 打印测试分隔线
    println("\n========== 路由分组测试 ==========\n")
    
    // ========== 1. 根分组（无前缀） ==========
    r.GET("/", func(c *gee.Context) {
        c.String(http.StatusOK, "根分组: 首页")
    })
    
    r.GET("/about", func(c *gee.Context) {
        c.String(http.StatusOK, "根分组: 关于页面")
    })
    
    // ========== 2. 一级分组 ==========
    api := r.Group("/api")
    {
        api.GET("/users", func(c *gee.Context) {
            c.String(http.StatusOK, "API分组: 用户列表")
        })
        
        api.GET("/posts", func(c *gee.Context) {
            c.String(http.StatusOK, "API分组: 文章列表")
        })
        
        api.POST("/users", func(c *gee.Context) {
            c.String(http.StatusOK, "API分组: 创建用户")
        })
    }
    
    // ========== 3. 二级分组（嵌套） ==========
    admin := r.Group("/admin")
    {
        admin.GET("/dashboard", func(c *gee.Context) {
            c.String(http.StatusOK, "Admin分组: 仪表板")
        })
        
        // 三级分组（更深嵌套）
        userMgmt := admin.Group("/users")
        {
            userMgmt.GET("/", func(c *gee.Context) {
                c.String(http.StatusOK, "Admin/Users分组: 用户管理列表")
            })
            
            userMgmt.GET("/:id", func(c *gee.Context) {
                id := c.Param("id")
                c.String(http.StatusOK, fmt.Sprintf("Admin/Users分组: 查看用户 %s", id))
            })
            
            userMgmt.DELETE("/:id", func(c *gee.Context) {
                id := c.Param("id")
                c.String(http.StatusOK, fmt.Sprintf("Admin/Users分组: 删除用户 %s", id))
            })
        }
    }
    
    // ========== 4. 多级嵌套分组 ==========
    v1 := r.Group("/api/v1")
    {
        v1.GET("/status", func(c *gee.Context) {
            c.String(http.StatusOK, "API v1分组: 状态查询")
        })
        
        v2 := v1.Group("/v2")
        {
            v2.GET("/data", func(c *gee.Context) {
                c.String(http.StatusOK, "API v1/v2分组: 获取数据")
            })
            
            v3 := v2.Group("/v3")
            {
                v3.GET("/info", func(c *gee.Context) {
                    c.String(http.StatusOK, "API v1/v2/v3分组: 详细信息")
                })
            }
        }
    }
    
    // ========== 5. 测试分组与动态路由结合 ==========
    product := r.Group("/product")
    {
        product.GET("/:category", func(c *gee.Context) {
            category := c.Param("category")
            c.String(http.StatusOK, "商品分组: 分类 %s 的商品列表", category)
        })
        
        detail := product.Group("/detail")
        {
            detail.GET("/:id", func(c *gee.Context) {
                id := c.Param("id")
                c.String(http.StatusOK, "商品分组: 商品 %s 的详细信息", id)
            })
        }
    }
    
    // 打印路由信息
    fmt.Println("服务器启动成功！")
    fmt.Println("\n可测试的URL：")
    fmt.Println("  GET  /                           -> 根分组")
    fmt.Println("  GET  /about                      -> 根分组")
    fmt.Println("  GET  /api/users                  -> API分组")
    fmt.Println("  GET  /api/posts                  -> API分组")
    fmt.Println("  POST /api/users                  -> API分组")
    fmt.Println("  GET  /admin/dashboard            -> Admin分组")
    fmt.Println("  GET  /admin/users/               -> Admin/Users分组")
    fmt.Println("  GET  /admin/users/123            -> Admin/Users分组 + 动态参数")
    fmt.Println("  DELETE /admin/users/123          -> Admin/Users分组 + 动态参数")
    fmt.Println("  GET  /api/v1/status              -> 多级嵌套分组")
    fmt.Println("  GET  /api/v1/v2/data             -> 多级嵌套分组")
    fmt.Println("  GET  /api/v1/v2/v3/info          -> 多级嵌套分组")
    fmt.Println("  GET  /product/electronics        -> 分组 + 动态路由")
    fmt.Println("  GET  /product/detail/100         -> 分组 + 动态路由")
    fmt.Println("\n服务器运行在 http://localhost:8080")
    
    r.Run(":8080")
}