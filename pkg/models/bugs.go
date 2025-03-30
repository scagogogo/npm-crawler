package models

// Bugs 表示 NPM 包的问题跟踪信息
//
// 包含了用于报告和跟踪包相关问题的链接地址，通常指向 GitHub Issues 或其他问题跟踪系统
//
// 主要字段说明:
//   - URL: 问题跟踪系统的链接地址
type Bugs struct {
	URL string `json:"url"` // 问题跟踪系统的链接地址
}
