package models

// User 表示 NPM 包的用户信息
//
// 在 NPM 生态系统中，User 通常用于表示包的维护者、发布者或贡献者
// 类似于 Author，但通常用于表示多个相关用户
//
// 主要字段说明:
//   - Name: 用户名称
//   - Email: 用户电子邮件地址
//   - URL: 用户网站或个人主页（可选）
type User struct {
	Name  string `json:"name"`  // 用户名称
	Email string `json:"email"` // 用户电子邮件地址
	URL   string `json:"url"`   // 用户网站或个人主页
}
