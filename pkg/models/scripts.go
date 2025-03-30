package models

// Script 表示 NPM 包的脚本命令定义
//
// 包含 NPM 包中定义的各种脚本命令，这些脚本可以通过 npm run [script-name] 来执行
// 最常用的脚本是 test 和 start，分别用于运行测试和启动项目
//
// 主要字段说明:
//   - Test: 测试脚本命令
//   - Start: 启动项目脚本命令
type Script struct {
	Test  string `json:"test"`  // 测试脚本命令
	Start string `json:"start"` // 启动项目脚本命令
}
