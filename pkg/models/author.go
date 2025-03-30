package models

// Author 表示 NPM 包的作者信息
//
// 包含了 NPM 包作者的基本信息，通常从 package.json 中解析得到
//
// 主要字段说明:
//   - Name: 作者名称
//   - Email: 作者电子邮件（可选）
//   - Url: 作者网站或个人主页（可选）
type Author struct {
	Name  string `json:"name"`  // 作者名称
	Email string `json:"email"` // 作者电子邮件
	Url   string `json:"url"`   // 作者网站或个人主页
}
