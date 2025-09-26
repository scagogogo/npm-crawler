import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'NPM Crawler',
  description: 'High-performance NPM Registry client with multi-mirror source and proxy support',
  lang: 'zh-CN',
  base: '/npm-crawler/',
  
  themeConfig: {
    logo: 'https://cdn.worldvectorlogo.com/logos/npm-2.svg',
    
    nav: [
      { text: '首页', link: '/' },
      { text: 'API文档', link: '/api/' },
      { text: 'GitHub', link: 'https://github.com/scagogogo/npm-crawler' }
    ],

    sidebar: {
      '/': [
        {
          text: '介绍',
          items: [
            { text: '什么是 NPM Crawler', link: '/' },
            { text: '快速开始', link: '/getting-started' },
            { text: '安装指南', link: '/installation' }
          ]
        },
        {
          text: 'API参考',
          items: [
            { text: 'API 概述', link: '/api/' },
            { text: '中文API文档', link: '/api/zh' },
            { text: 'English API Docs', link: '/api/en' }
          ]
        },
        {
          text: '示例',
          items: [
            { text: '基本用法', link: '/examples/basic' },
            { text: '高级用法', link: '/examples/advanced' },
            { text: '镜像源配置', link: '/examples/mirrors' }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/npm-crawler' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2023-present NPM Crawler Team'
    },

    editLink: {
      pattern: 'https://github.com/scagogogo/npm-crawler/edit/main/docs/:path',
      text: '在 GitHub 上编辑此页面'
    },

    lastUpdated: {
      text: '最后更新于',
      formatOptions: {
        dateStyle: 'short',
        timeStyle: 'medium'
      }
    },

    docFooter: {
      prev: '上一页',
      next: '下一页'
    },

    outline: {
      label: '页面导航'
    },

    returnToTopLabel: '回到顶部',
    sidebarMenuLabel: '菜单',
    darkModeSwitchLabel: '主题',
    lightModeSwitchTitle: '切换到浅色模式',
    darkModeSwitchTitle: '切换到深色模式'
  },

  markdown: {
    theme: {
      light: 'github-light',
      dark: 'github-dark'
    },
    lineNumbers: true
  }
})
