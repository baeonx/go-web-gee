package gee

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	// 路由参数
    Params map[string]string
    // 响应信息
    StatusCode int
}

// Param 获取路由参数
func (c *Context) Param(key string) string {
    if c.Params == nil {
        return ""
    }
    return c.Params[key]
}

// Query 获取查询参数
func (c *Context) Query(key string) string {
    return c.Req.URL.Query().Get(key)
}

// String 返回文本响应
func (c *Context) String(statusCode int, format string, args ...interface{}) {
    c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
    c.Writer.WriteHeader(statusCode)
    fmt.Fprintf(c.Writer, format, args...)
}

// JSON 返回JSON响应（简化版）
func (c *Context) JSON(statusCode int, data interface{}) {
    c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
    c.Writer.WriteHeader(statusCode)
    
    // 使用 json.NewEncoder 而不是 fmt.Fprintf
    encoder := json.NewEncoder(c.Writer)
    if err := encoder.Encode(data); err != nil {
        // 如果编码失败，返回错误信息
        http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
    }
}

// HTML 返回HTML响应
func (c *Context) HTML(statusCode int, format string, args ...interface{}) {
    c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
    c.Writer.WriteHeader(statusCode)
    fmt.Fprintf(c.Writer, format, args...)
}


