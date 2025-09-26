# 中文 API 文档

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/npm-crawler.svg)](https://pkg.go.dev/github.com/scagogogo/npm-crawler)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

_完整的 API 参考和使用指南_

</div>

## 目录

- [概述](#概述)
- [安装](#安装)
- [快速开始](#快速开始)
- [核心 API 参考](#核心-api-参考)
  - [Registry 客户端](#registry-客户端)
  - [配置选项](#配置选项)
  - [镜像源](#镜像源)
- [数据模型](#数据模型)
- [高级用法](#高级用法)
- [错误处理](#错误处理)
- [最佳实践](#最佳实践)
- [故障排除](#故障排除)
- [示例](#示例)

## 概述

NPM Crawler 提供了一个全面的 Go 客户端用于与 NPM 注册表交互。本文档涵盖了所有可用的 API、使用模式和将库集成到应用程序中的最佳实践。

### 主要特性

- **多注册表支持**: 官方 NPM、淘宝、华为云等多种镜像源
- **代理配置**: 支持企业环境的 HTTP 代理
- **上下文支持**: 取消和超时处理
- **类型安全**: 完整的 NPM 元数据 Go 类型定义
- **高性能**: 针对并发访问优化

## 安装

```bash
go get github.com/scagogogo/npm-crawler
```

**要求:**
- Go 1.20 或更高版本
- 网络连接（或访问 NPM 注册表镜像）

## 快速开始

### 基本包信息获取

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/scagogogo/npm-crawler/pkg/registry"
)

func main() {
    // 使用默认设置创建客户端（官方 NPM 注册表）
    client := registry.NewRegistry()
    
    // 设置超时上下文
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // 获取包信息
    pkg, err := client.GetPackageInformation(ctx, "react")
    if err != nil {
        log.Fatalf("获取包信息失败: %v", err)
    }
    
    fmt.Printf("包名: %s\n", pkg.Name)
    fmt.Printf("描述: %s\n", pkg.Description)
    fmt.Printf("最新版本: %s\n", pkg.DistTags["latest"])
    fmt.Printf("许可证: %s\n", pkg.License)
    fmt.Printf("主页: %s\n", pkg.Homepage)
}
```

## 核心 API 参考

### Registry 客户端

#### 创建客户端

##### `NewRegistry(options ...*Options) *Registry`

创建一个新的 Registry 客户端实例。

**参数:**
- `options` - 可选的配置选项

**示例:**
```go
// 使用默认配置
client := registry.NewRegistry()

// 使用自定义配置
options := registry.NewOptions().
    SetRegistryURL("https://registry.npmjs.org").
    SetProxy("http://proxy.example.com:8080")
client := registry.NewRegistry(options)
```

#### 核心方法

##### `GetPackageInformation(ctx context.Context, packageName string) (*models.Package, error)`

获取指定 NPM 包的详细信息。

**参数:**
- `ctx` - 上下文，用于取消和超时控制
- `packageName` - 要查询的包名称

**返回值:**
- `*models.Package` - 完整的包信息
- `error` - 如果请求失败则返回错误

**示例:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

pkg, err := client.GetPackageInformation(ctx, "lodash")
if err != nil {
    return fmt.Errorf("获取包信息失败: %w", err)
}

// 访问包数据
fmt.Printf("名称: %s\n", pkg.Name)
fmt.Printf("最新版本: %s\n", pkg.DistTags["latest"])
fmt.Printf("作者: %s\n", pkg.Author.Name)

// 访问特定版本
if version, exists := pkg.Versions["4.17.21"]; exists {
    fmt.Printf("版本 4.17.21 依赖: %+v\n", version.Dependencies)
}
```

##### `GetRegistryInformation(ctx context.Context) (*models.RegistryInformation, error)`

获取 NPM 注册表的状态和元数据信息。

**参数:**
- `ctx` - 上下文，用于取消和超时控制

**返回值:**
- `*models.RegistryInformation` - 注册表状态信息
- `error` - 如果请求失败则返回错误

**示例:**
```go
info, err := client.GetRegistryInformation(ctx)
if err != nil {
    return fmt.Errorf("获取注册表信息失败: %w", err)
}

fmt.Printf("注册表: %s\n", info.DbName)
fmt.Printf("包总数: %d\n", info.DocCount)
fmt.Printf("数据库大小: %d 字节\n", info.DataSize)
fmt.Printf("磁盘使用: %d 字节\n", info.DiskSize)
```

##### `SearchPackages(ctx context.Context, query string, limit int) (*models.SearchResult, error)`

搜索 NPM 包。

**参数:**
- `ctx` - 上下文
- `query` - 搜索关键字
- `limit` - 返回结果数量限制，默认为 20

**返回值:**
- `*models.SearchResult` - 搜索结果
- `error` - 如果请求失败则返回错误

**示例:**
```go
result, err := client.SearchPackages(ctx, "react", 10)
if err != nil {
    return fmt.Errorf("搜索失败: %w", err)
}

fmt.Printf("找到 %d 个结果\n", result.Total)
for _, obj := range result.Objects {
    pkg := obj.Package
    fmt.Printf("包名: %s\n", pkg.Name)
    fmt.Printf("版本: %s\n", pkg.Version)
    fmt.Printf("描述: %s\n", pkg.Description)
    fmt.Printf("评分: %.2f\n", obj.Score.Final)
    fmt.Println("---")
}
```

##### `GetPackageVersion(ctx context.Context, packageName, version string) (*models.Version, error)`

获取指定包的特定版本信息。

**参数:**
- `ctx` - 上下文
- `packageName` - 包名称
- `version` - 版本号或标签（如 "1.0.0" 或 "latest"）

**返回值:**
- `*models.Version` - 版本详细信息
- `error` - 如果请求失败则返回错误

**示例:**
```go
version, err := client.GetPackageVersion(ctx, "react", "18.2.0")
if err != nil {
    return fmt.Errorf("获取版本信息失败: %w", err)
}

fmt.Printf("版本: %s\n", version.Version)
fmt.Printf("描述: %s\n", version.Description)
fmt.Printf("依赖: %+v\n", version.Dependencies)
fmt.Printf("开发依赖: %+v\n", version.DevDependencies)
```

##### `GetDownloadStats(ctx context.Context, packageName, period string) (*models.DownloadStats, error)`

获取指定包的下载统计信息。

**参数:**
- `ctx` - 上下文
- `packageName` - 包名称
- `period` - 统计周期（"last-day", "last-week", "last-month"）

**返回值:**
- `*models.DownloadStats` - 下载统计信息
- `error` - 如果请求失败则返回错误

**示例:**
```go
stats, err := client.GetDownloadStats(ctx, "react", "last-week")
if err != nil {
    return fmt.Errorf("获取下载统计失败: %w", err)
}

fmt.Printf("包: %s\n", stats.Package)
fmt.Printf("下载次数: %d\n", stats.Downloads)
fmt.Printf("统计周期: %s 到 %s\n", stats.Start, stats.End)
```

##### `GetOptions() *Options`

返回当前注册表客户端的配置选项。

**返回值:**
- `*Options` - 当前配置选项

### 配置选项

#### Options 结构

```go
type Options struct {
    RegistryURL string  // NPM 仓库服务器的 URL 地址
    Proxy       string  // HTTP 代理服务器的 URL
}
```

#### 方法

##### `NewOptions() *Options`

创建并返回一个新的默认配置选项实例。

**默认配置:**
- RegistryURL: "https://registry.npmjs.org"
- Proxy: 无代理设置

##### `SetRegistryURL(url string) *Options`

设置 NPM 仓库服务器的 URL 地址。

**参数:**
- `url` - NPM 仓库 URL 地址

**返回值:**
- `*Options` - 更新后的选项对象（支持链式调用）

##### `SetProxy(proxyUrl string) *Options`

设置 HTTP 代理服务器的 URL 地址。

**参数:**
- `proxyUrl` - HTTP 代理服务器的 URL 地址

**返回值:**
- `*Options` - 更新后的选项对象（支持链式调用）

**示例:**
```go
options := registry.NewOptions().
    SetRegistryURL("https://registry.npmmirror.com").
    SetProxy("http://proxy.corp.com:8080")
```

##### `GetHttpClient() (*http.Client, error)`

根据当前选项配置创建并返回一个 HTTP 客户端。

**返回值:**
- `*http.Client` - 配置好的 HTTP 客户端
- `error` - 如果代理 URL 解析失败

### 镜像源

NPM Crawler 内置支持多种镜像源，特别适合中国大陆用户：

#### 官方镜像

```go
// NPM 官方注册表 (全球)
client := registry.NewRegistry()

// Yarn 官方镜像 (全球)
client := registry.NewYarnRegistry()
```

#### 中国镜像源

```go
// 淘宝 NPM 镜像 (中国)
client := registry.NewTaoBaoRegistry()

// NPM Mirror (中国)
client := registry.NewNpmMirrorRegistry()

// 华为云镜像 (中国)
client := registry.NewHuaWeiCloudRegistry()

// 腾讯云镜像 (中国)
client := registry.NewTencentRegistry()

// CNPM 镜像 (中国)
client := registry.NewCnpmRegistry()

// NPM CouchDB 镜像
client := registry.NewNpmjsComRegistry()
```

## 数据模型

### Package 模型

表示 NPM 包的完整信息：

```go
type Package struct {
    ID             string                 `json:"_id"`             // 包 ID
    Name           string                 `json:"name"`            // 包名称
    Description    string                 `json:"description"`     // 包描述
    DistTags       map[string]string      `json:"dist-tags"`       // 分发标签
    Versions       map[string]Version     `json:"versions"`        // 版本信息
    Maintainers    []Maintainer           `json:"maintainers"`     // 维护者
    Time           map[string]string      `json:"time"`            // 时间信息
    Repository     Repository             `json:"repository"`      // 仓库信息
    Homepage       string                 `json:"homepage"`        // 主页
    License        string                 `json:"license"`         // 许可证
    Keywords       []string               `json:"keywords"`        // 关键词
    Author         Author                 `json:"author"`          // 作者
    // ... 其他字段
}
```

**常用操作:**
```go
// 获取最新版本
latestVersion := pkg.DistTags["latest"]

// 列出所有可用版本
for version := range pkg.Versions {
    fmt.Printf("可用版本: %s\n", version)
}

// 获取特定版本详情
if versionInfo, exists := pkg.Versions["1.0.0"]; exists {
    fmt.Printf("依赖: %+v\n", versionInfo.Dependencies)
    fmt.Printf("开发依赖: %+v\n", versionInfo.DevDependencies)
}
```

### Version 模型

```go
type Version struct {
    Name            string               `json:"name"`            // 包名称
    Version         string               `json:"version"`         // 版本号
    Description     string               `json:"description"`     // 描述
    Main            string               `json:"main"`            // 入口点
    Scripts         *Script              `json:"scripts"`         // NPM 脚本
    Dependencies    map[string]string    `json:"dependencies"`    // 运行时依赖
    DevDependencies map[string]string    `json:"devDependencies"` // 开发依赖
    Repository      *Repository          `json:"repository"`      // 仓库
    License         string               `json:"license"`         // 许可证
    Dist            *Dist                `json:"dist"`            // 分发信息
    // ... 其他字段
}
```

### SearchResult 模型

```go
type SearchResult struct {
    Objects []SearchObject `json:"objects"` // 搜索结果对象
    Total   int            `json:"total"`   // 总匹配数量
    Time    string         `json:"time"`    // 搜索耗时
}

type SearchObject struct {
    Package     SearchPackage `json:"package"`     // 包信息
    Score       Score         `json:"score"`       // 评分
    SearchScore float64       `json:"searchScore"` // 搜索得分
}
```

### DownloadStats 模型

```go
type DownloadStats struct {
    Downloads int    `json:"downloads"` // 下载次数
    Start     string `json:"start"`     // 统计开始日期
    End       string `json:"end"`       // 统计结束日期
    Package   string `json:"package"`   // 包名称
}
```

### RegistryInformation 模型

```go
type RegistryInformation struct {
    DbName            string `json:"db_name"`              // 数据库名称
    DocCount          int    `json:"doc_count"`            // 包总数
    DocDelCount       int    `json:"doc_del_count"`        // 已删除包数
    UpdateSeq         int    `json:"update_seq"`           // 更新序列
    CompactRunning    bool   `json:"compact_running"`      // 压缩状态
    DiskSize          int64  `json:"disk_size"`            // 磁盘使用
    DataSize          int64  `json:"data_size"`            // 数据大小
    InstanceStartTime string `json:"instance_start_time"`  // 启动时间
    // ... 其他字段
}
```

## 高级用法

### 并发包查询

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/scagogogo/npm-crawler/pkg/registry"
)

func main() {
    client := registry.NewRegistry()
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    packages := []string{"react", "vue", "angular", "lodash", "express"}
    results := make(chan string, len(packages))
    
    var wg sync.WaitGroup
    
    for _, pkgName := range packages {
        wg.Add(1)
        go func(name string) {
            defer wg.Done()
            
            pkg, err := client.GetPackageInformation(ctx, name)
            if err != nil {
                results <- fmt.Sprintf("%s: 错误 - %v", name, err)
                return
            }
            
            results <- fmt.Sprintf("%s: %s", pkg.Name, pkg.DistTags["latest"])
        }(pkgName)
    }
    
    wg.Wait()
    close(results)
    
    for result := range results {
        fmt.Println(result)
    }
}
```

### 代理配置

```go
// 基本代理
options := registry.NewOptions().SetProxy("http://proxy.corp.com:8080")

// 带认证的代理
options.SetProxy("http://username:password@proxy.corp.com:8080")

// 使用环境变量
import "os"
if proxy := os.Getenv("HTTP_PROXY"); proxy != "" {
    options.SetProxy(proxy)
}
```

### 错误处理和重试

```go
import (
    "context"
    "errors"
    "time"
)

func getPackageWithRetry(client *registry.Registry, packageName string, maxRetries int) (*models.Package, error) {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        pkg, err := client.GetPackageInformation(ctx, packageName)
        cancel()
        
        if err == nil {
            return pkg, nil
        }
        
        lastErr = err
        
        // 对于某些错误类型，不需要重试
        if errors.Is(err, context.Canceled) {
            break
        }
        
        // 等待后重试
        if i < maxRetries-1 {
            time.Sleep(time.Duration(i+1) * time.Second)
        }
    }
    
    return nil, fmt.Errorf("重试 %d 次后仍然失败: %w", maxRetries, lastErr)
}
```

### 批量操作

```go
func getMultiplePackages(client *registry.Registry, packageNames []string) (map[string]*models.Package, error) {
    ctx := context.Background()
    results := make(map[string]*models.Package)
    errors := make(map[string]error)
    
    type result struct {
        name string
        pkg  *models.Package
        err  error
    }
    
    ch := make(chan result, len(packageNames))
    
    // 启动所有查询
    for _, name := range packageNames {
        go func(packageName string) {
            pkg, err := client.GetPackageInformation(ctx, packageName)
            ch <- result{packageName, pkg, err}
        }(name)
    }
    
    // 收集结果
    for i := 0; i < len(packageNames); i++ {
        res := <-ch
        if res.err != nil {
            errors[res.name] = res.err
        } else {
            results[res.name] = res.pkg
        }
    }
    
    if len(errors) > 0 {
        return results, fmt.Errorf("部分包获取失败: %+v", errors)
    }
    
    return results, nil
}
```

## 错误处理

### 常见错误类型

```go
import (
    "context"
    "errors"
    "net"
    "net/http"
)

func handleError(err error) {
    switch {
    case errors.Is(err, context.DeadlineExceeded):
        fmt.Println("请求超时")
    case errors.Is(err, context.Canceled):
        fmt.Println("请求被取消")
    default:
        // 检查 HTTP 状态码
        var httpErr *http.Response
        if errors.As(err, &httpErr) {
            switch httpErr.StatusCode {
            case 404:
                fmt.Println("包不存在")
            case 500:
                fmt.Println("服务器内部错误")
            default:
                fmt.Printf("HTTP 错误: %d\n", httpErr.StatusCode)
            }
        }
        
        // 检查网络错误
        var netErr net.Error
        if errors.As(err, &netErr) {
            if netErr.Timeout() {
                fmt.Println("网络超时")
            } else {
                fmt.Println("网络错误:", netErr)
            }
        }
    }
}
```

## 最佳实践

### 1. 上下文管理

```go
// 为每个请求设置合理的超时
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()

// 支持取消操作
ctx, cancel := context.WithCancel(context.Background())
go func() {
    // 某些条件下取消请求
    time.Sleep(5 * time.Second)
    cancel()
}()
```

### 2. 镜像选择

```go
func selectBestMirror() *registry.Registry {
    // 根据地理位置选择最佳镜像
    if isInChina() {
        return registry.NewNpmMirrorRegistry() // 中国用户优选
    }
    return registry.NewRegistry() // 国际用户使用官方源
}

func isInChina() bool {
    // 实现地理位置检测逻辑
    return false
}
```

### 3. 缓存策略

```go
import (
    "sync"
    "time"
)

type PackageCache struct {
    cache map[string]*CacheEntry
    mutex sync.RWMutex
}

type CacheEntry struct {
    Package   *models.Package
    Timestamp time.Time
    TTL       time.Duration
}

func (c *PackageCache) Get(packageName string) (*models.Package, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    entry, exists := c.cache[packageName]
    if !exists {
        return nil, false
    }
    
    if time.Since(entry.Timestamp) > entry.TTL {
        return nil, false // 过期
    }
    
    return entry.Package, true
}

func (c *PackageCache) Set(packageName string, pkg *models.Package, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.cache[packageName] = &CacheEntry{
        Package:   pkg,
        Timestamp: time.Now(),
        TTL:       ttl,
    }
}
```

### 4. 性能优化

```go
// 使用连接池
import "net/http"

func createOptimizedClient() *registry.Registry {
    transport := &http.Transport{
        MaxIdleConns:          100,
        MaxIdleConnsPerHost:   10,
        IdleConnTimeout:       90 * time.Second,
        TLSHandshakeTimeout:   10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }
    
    httpClient := &http.Client{
        Transport: transport,
        Timeout:   30 * time.Second,
    }
    
    // 注意：当前版本不支持自定义 HTTP 客户端
    // 这是一个示例，展示了可能的优化方向
    return registry.NewRegistry()
}
```

## 故障排除

### 1. 网络连接问题

**问题**: 请求超时或连接失败

**解决方案**:
```go
// 增加超时时间
ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

// 使用更快的镜像
client := registry.NewNpmMirrorRegistry() // 对于中国用户

// 检查代理配置
options := registry.NewOptions().SetProxy("") // 禁用代理
client := registry.NewRegistry(options)
```

### 2. 包不存在 (404)

**问题**: 包存在但返回 404

**解决方案**:
```go
// 检查包名拼写
pkg, err := client.GetPackageInformation(ctx, "react") // 不是 "React"

// 尝试不同的注册表
client := registry.NewYarnRegistry()

// 验证包是否在 npmjs.com 上存在
```

### 3. 代理问题

**问题**: 企业代理阻止请求

**解决方案**:
```go
// 使用认证代理
options := registry.NewOptions().
    SetProxy("http://username:password@proxy.corp.com:8080")

// 尝试不同的代理格式
options.SetProxy("http://proxy.corp.com:8080")

// 为特定域名绕过代理（需要自定义 HTTP 客户端）
```

### 4. 速率限制

**问题**: 过多请求导致速率限制

**解决方案**:
```go
// 在请求之间添加延迟
time.Sleep(100 * time.Millisecond)

// 使用不同镜像进行负载分布
registries := []*registry.Registry{
    registry.NewRegistry(),
    registry.NewNpmMirrorRegistry(),
    registry.NewTaoBaoRegistry(),
}

// 实现请求队列
type RequestQueue struct {
    ch chan func()
}

func (rq *RequestQueue) Execute(fn func()) {
    rq.ch <- fn
}
```

### 5. 大包内存使用

**问题**: 查询大包时内存使用过高

**解决方案**:
```go
// 只访问需要的字段
pkg, err := client.GetPackageInformation(ctx, packageName)
if err != nil {
    return err
}

// 立即提取需要的数据
result := struct {
    Name        string
    Description string
    Latest      string
}{
    Name:        pkg.Name,
    Description: pkg.Description,
    Latest:      pkg.DistTags["latest"],
}

// 清除原始包数据
pkg = nil
```

## 示例

完整的示例代码可以在以下位置找到：

- [基本用法示例](/examples/basic)
- [高级用法示例](/examples/advanced)
- [镜像源配置示例](/examples/mirrors)
- [GitHub 仓库示例](https://github.com/scagogogo/npm-crawler/tree/main/examples)

---

## 贡献

我们欢迎对此文档的改进建议！请随时：

- 报告问题或不一致之处
- 建议新的示例或用例
- 添加翻译
- 改进现有内容

更多信息请参阅我们的 [贡献指南](https://github.com/scagogogo/npm-crawler/blob/main/README.md#contribution-guide)。

## 许可证

此文档是 NPM Crawler 项目的一部分，采用 MIT 许可证。详见 [LICENSE](https://github.com/scagogogo/npm-crawler/blob/main/LICENSE)。
