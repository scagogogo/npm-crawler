# NPM Crawler 文档

这是 NPM Crawler 项目的文档站点，使用 VitePress 构建。

## 本地开发

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

## 文档结构

- `index.md` - 首页
- `getting-started.md` - 快速开始指南
- `installation.md` - 安装指南
- `api/` - API 文档
  - `index.md` - API 概述
  - `zh.md` - 中文 API 文档
  - `en.md` - 英文 API 文档
- `examples/` - 示例代码
  - `basic.md` - 基本用法示例
  - `advanced.md` - 高级用法示例
  - `mirrors.md` - 镜像源配置示例

## 自动部署

文档通过 GitHub Actions 自动部署到 GitHub Pages。每当有推送到 `main` 分支且影响 `docs/` 目录的更改时，会自动触发部署。

## 贡献

欢迎为文档做出贡献！请确保：

1. 遵循现有的文档风格
2. 在本地测试构建是否成功
3. 提供清晰的提交信息
