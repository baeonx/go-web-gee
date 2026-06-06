# Gee - 从零实现的 Go Web 框架
Gee 是一个**学习性质**的 Go Web 框架，通过 7 天迭代从零构建，深入理解 Web 框架的核心原理。API 设计参考 Gin 框架。

> 📖 **配套教程**：[7天用Go从零实现Web框架Gee](https://geektutu.com/post/gee.html)

## 🚀 快速体验


```bash
# 克隆项目
git clone https://github.com/baeonx/go-web-gee.git

# 进入对应天的示例目录
cd go-web-gee/day1

# 运行示例
go run main.go
```
📅 开发路线图
Day 1 - 框架雏形 & 路由映射
Day 2 - 上下文 Context & 路由分离
Day 3 - 动态路由 (Trie 树)
Day 4 - 路由分组控制
Day 5 - 中间件链
Day 6 - 模板引擎支持
Day 7 - 错误恢复机制


Day 1：框架雏形 & 路由映射
📝 实现功能
✅ 实现 http.Handler 接口，接管所有 HTTP 请求
✅ 设计路由映射表，支持 GET/POST 方法
✅ 封装启动函数 Run()
✅ 路由路径自动规范化（添加 / 前缀）

📁 代码结构
text
day1/
├── gee/
│   └── gee.go          # 框架核心（~90行）
├── main.go             # 使用示例
└── go.mod              # 模块依赖

