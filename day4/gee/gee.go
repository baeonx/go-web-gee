// Package gee 提供了一个轻量级的 HTTP Web 框架。
// 它支持路由注册、请求分发等核心功能。
package gee

import (
    "net/http"
)


type RouterGroup struct{
    prefix string
    parent *RouterGroup
    engine *Engine
}

// HandlerFunc 定义了框架处理 HTTP 请求的函数类型。
// 该签名与标准库 http.HandlerFunc 兼容。
type HandlerFunc func(c *Context)

// Engine 是框架的核心引擎，负责管理路由和处理 HTTP 请求。
// 它实现了 http.Handler 接口，可直接传递给 http.ListenAndServe。
type Engine struct {
    // router 是路由映射表，键为 "METHOD-path" 格式，值为对应的处理函数
    router *Router
    *RouterGroup
    groups []*RouterGroup
}

// New 创建并初始化一个新的 Engine 实例。
// 返回的 Engine 指针可用于注册路由和启动 HTTP 服务。
func New() *Engine {
    engine := &Engine{
        router: newRouter(),
        groups: make([]*RouterGroup, 0),
    }
    
    // 创建根分组
    rootGroup := &RouterGroup{
        prefix: "",
        parent: nil,
        engine: engine,  // 根分组指向 engine
    }
    
    // 将根分组设置为 Engine 的嵌入字段
    engine.RouterGroup = rootGroup
    engine.groups = append(engine.groups, rootGroup)
    
    return engine
}

// normalizePattern 规范化路由路径格式，确保所有路径都以 '/' 开头。
// 参数 pattern: 原始路由模式
// 返回值: 规范化后的路径
func (engine *Engine) normalizePattern(pattern string) string {
    if len(pattern) == 0 {
        return "/"
    }
    if pattern[0] != '/' {
        return "/" + pattern
    }
    return pattern
}


func (group *RouterGroup) Group(pattern string) *RouterGroup {
    normalizedPrefix := group.engine.normalizePattern(pattern)
    newGroup := &RouterGroup{
        prefix: group.prefix + normalizedPrefix,
        parent: group,
        engine: group.engine,  
    }
    group.engine.groups = append(group.engine.groups, newGroup)
    return newGroup
}


func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
    group.engine.router.addRoute(method, pattern, handler)
}

// GET 注册一个处理 HTTP GET 请求的路由。
// 参数 pattern: 路由路径，例如 "/hello"、"/user/profile"
// 参数 handler: 处理该请求的函数，必须符合 HandlerFunc 类型定义
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
    pattern = group.prefix + group.engine.normalizePattern(pattern)
    group.addRoute("GET", pattern, handler)
}

// POST 注册一个处理 HTTP POST 请求的路由。
// 参数 pattern: 路由路径，例如 "/user/create"、"/api/login"
// 参数 handler: 处理该请求的函数，必须符合 HandlerFunc 类型定义
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
    pattern = group.prefix +group.engine.normalizePattern(pattern)
    group.addRoute("POST", pattern, handler)
}

// PUT 注册一个处理 HTTP PUT 请求的路由
func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	pattern = group.prefix + group.engine.normalizePattern(pattern)
	group.addRoute("PUT", pattern, handler)
}

// DELETE 注册一个处理 HTTP DELETE 请求的路由
func (group *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	pattern = group.prefix +group.engine.normalizePattern(pattern)
	group.addRoute("DELETE", pattern, handler)
}

// Run 启动 HTTP 服务器并监听指定的端口。
// 这是一个阻塞调用，服务器会持续运行直到程序退出或发生错误。
// 参数 port: 服务器监听的端口，格式为 ":8080" 或 "127.0.0.1:8080"
func (engine *Engine) Run(port string) {
    http.ListenAndServe(port, engine)
}

// ServeHTTP 实现 http.Handler 接口的核心方法。
// 它根据请求的方法和路径从路由表中查找匹配的处理函数并执行。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    c := engine.newContext(w, req)
    // 使用 router 的 handle 方法处理请求
	engine.router.handle(c)
}


func (engine *Engine) newContext(w http.ResponseWriter, req *http.Request) *Context {
    return &Context{
        Writer: w,
        Req:    req,
        Path:   req.URL.Path,
        Method: req.Method,
        StatusCode: http.StatusOK,  // 默认 200 OK
        Params:     make(map[string]string), // 初始化 Params map
    }
}