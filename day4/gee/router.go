package gee

import (
    "strings"
    "net/http"
)

// Router 负责路由管理
type Router struct {
    roots    map[string]*node       // 每个 HTTP 方法一棵 Trie 树
    handlers map[string]HandlerFunc // 存储路由对应的处理函数
}

// NewRouter 创建新的 Router 实例
func newRouter() *Router {
    return &Router{
        roots:    make(map[string]*node),
        handlers: make(map[string]HandlerFunc),
    }
}

// parsePattern 将路径切分成片段 
// 例如："/user/:id/post" -> ["user", ":id", "post"]
func parsePattern(pattern string) []string {
    vs := strings.Split(pattern, "/")
    parts := make([]string, 0)
    for _, item := range vs {
        if item != "" {
            parts = append(parts, item)
            if item[0] == '*' {
                break  // * 通配符必须放在最后，遇到就停止
            }
        }
    }
    return parts
}

// addRoute 添加路由到 Trie 树
func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
    // 获取该方法的根节点
    if _, ok := r.roots[method]; !ok {
        r.roots[method] = &node{}
    }
    
	
    parts := parsePattern(pattern)
    key := method + "-" + pattern
    
    // 插入到 Trie 树
    r.roots[method].insert(pattern, parts, 0)
    r.handlers[key] = handler
}

// getRoute 查找路由，返回节点和解析出的参数
func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
    root, ok := r.roots[method]
    if !ok {
        return nil, nil
    }
    
    searchParts := parsePattern(path)
    n := root.search(searchParts, 0)
    
    if n == nil {
        return nil, nil
    }
    
    // 解析参数
    params := make(map[string]string)
    parts := parsePattern(n.pattern)
    for index, part := range parts {
        if part[0] == ':' {
            params[part[1:]] = searchParts[index]
        }
        if part[0] == '*' && len(part) > 1 {
            params[part[1:]] = strings.Join(searchParts[index:], "/")
            break
        }
    }
    return n, params
}

// handle 处理 HTTP 请求
func (r *Router) handle(c *Context) {
    n, params := r.getRoute(c.Method, c.Path)
    if n != nil {
        key := c.Method + "-" + n.pattern
        c.Params = params  // 需要给 Context 添加 Params 字段
        r.handlers[key](c)
    } else {
        c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
    }
}