package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 完整路由路径，如 /p/:lang/doc
	part     string  // 当前节点的片段，如 :lang 或 doc
	children []*node // 子节点列表
	isWild   bool    // 是否为通配符节点（: 或 * 开头）
}

// insert 插入路由到 Trie 树
func (n *node) insert(pattern string, parts []string, height int) {
	// 到达末尾，设置 pattern
	if height == len(parts) {
		n.pattern = pattern
		return
	}

	part := parts[height]

	// 查找匹配的子节点（使用 matchChild 方法）
	child := n.matchChild(part)
	if child == nil {
		// 创建新节点
		child = &node{
			part:     part,
			children: []*node{},
			isWild:   part[0] == ':' || part[0] == '*',
			pattern:  "",
		}
		n.children = append(n.children, child)
	}

	// 递归插入
	child.insert(pattern, parts, height+1)
}

// search 查找匹配的路由节点
func (n *node) search(parts []string, height int) *node {
	// 到达路径末尾 或 当前节点是通配符*（*匹配剩余所有）
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]

	// 查找所有匹配的子节点（使用 matchChildren 方法）
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// matchChild 查找第一个匹配的子节点（用于插入）
// 注意：插入时只匹配静态节点，不匹配动态节点，避免重复创建
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part && !child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 查找所有匹配的子节点（用于查找）
// 返回所有静态匹配和动态匹配的子节点
func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}