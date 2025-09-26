package registry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 注意：这些测试需要网络连接，可能因为网络或服务可用性而失败
// 在实际的CI/CD环境中，建议使用mock或集成测试环境

func TestSearchPackages(t *testing.T) {
	t.Skip("跳过网络依赖的测试，避免CI/CD环境中的不稳定性")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试搜索react包
	result, err := registry.SearchPackages(ctx, "react", 5)
	if err != nil {
		t.Logf("搜索包时出错（可能是网络问题）: %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.True(t, len(result.Objects) > 0, "搜索结果应该包含至少一个包")
	assert.True(t, result.Total > 0, "总数应该大于0")

	// 检查第一个结果是否包含react
	if len(result.Objects) > 0 {
		firstPkg := result.Objects[0]
		assert.NotEmpty(t, firstPkg.Package.Name, "包名不应为空")
		assert.NotEmpty(t, firstPkg.Package.Version, "版本不应为空")
	}
}

func TestSearchPackagesWithLimit(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试限制结果数量
	result, err := registry.SearchPackages(ctx, "lodash", 3)
	if err != nil {
		t.Logf("搜索包时出错: %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.True(t, len(result.Objects) <= 3, "结果数量应该不超过限制")
}

func TestSearchPackagesDefaultLimit(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试默认限制（limit=0应该使用默认值20）
	result, err := registry.SearchPackages(ctx, "express", 0)
	if err != nil {
		t.Logf("搜索包时出错: %v", err)
		return
	}

	assert.NotNil(t, result)
	// 默认limit应该是20，但实际结果可能少于20
	assert.True(t, len(result.Objects) <= 20, "结果数量应该不超过默认限制20")
}

func TestGetPackageVersion(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试获取特定版本的包信息
	version, err := registry.GetPackageVersion(ctx, "lodash", "4.17.21")
	if err != nil {
		t.Logf("获取包版本时出错: %v", err)
		return
	}

	assert.NotNil(t, version)
	assert.Equal(t, "lodash", version.Name, "包名应该匹配")
	assert.Equal(t, "4.17.21", version.Version, "版本号应该匹配")
	assert.NotEmpty(t, version.Description, "描述不应为空")
}

func TestGetPackageVersionLatest(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试获取latest版本
	version, err := registry.GetPackageVersion(ctx, "react", "latest")
	if err != nil {
		t.Logf("获取最新版本时出错: %v", err)
		return
	}

	assert.NotNil(t, version)
	assert.Equal(t, "react", version.Name, "包名应该匹配")
	assert.NotEmpty(t, version.Version, "版本号不应为空")
	assert.NotEmpty(t, version.Description, "描述不应为空")
}

func TestGetDownloadStats(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试获取最近一周的下载统计
	stats, err := registry.GetDownloadStats(ctx, "react", "last-week")
	if err != nil {
		t.Logf("获取下载统计时出错: %v", err)
		return
	}

	assert.NotNil(t, stats)
	assert.Equal(t, "react", stats.Package, "包名应该匹配")
	assert.True(t, stats.Downloads >= 0, "下载次数应该大于等于0")
	assert.NotEmpty(t, stats.Start, "开始日期不应为空")
	assert.NotEmpty(t, stats.End, "结束日期不应为空")
}

func TestGetDownloadStatsLastDay(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试获取最近一天的下载统计
	stats, err := registry.GetDownloadStats(ctx, "lodash", "last-day")
	if err != nil {
		t.Logf("获取昨日下载统计时出错: %v", err)
		return
	}

	assert.NotNil(t, stats)
	assert.Equal(t, "lodash", stats.Package)
	assert.True(t, stats.Downloads >= 0)
}

func TestGetDownloadStatsLastMonth(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	registry := NewRegistry()
	ctx := context.Background()

	// 测试获取最近一个月的下载统计
	stats, err := registry.GetDownloadStats(ctx, "express", "last-month")
	if err != nil {
		t.Logf("获取月度下载统计时出错: %v", err)
		return
	}

	assert.NotNil(t, stats)
	assert.Equal(t, "express", stats.Package)
	assert.True(t, stats.Downloads >= 0)
}

// 测试各种镜像源是否支持新的API
func TestNewAPIWithDifferentMirrors(t *testing.T) {
	t.Skip("跳过网络依赖的测试")

	ctx := context.Background()

	mirrors := []struct {
		name     string
		registry *Registry
	}{
		{"Official", NewRegistry()},
		{"CNPM", NewCnpmRegistry()},
		{"Yarn", NewYarnRegistry()},
		// 注意：搜索API可能只在官方registry支持
	}

	for _, mirror := range mirrors {
		t.Run(mirror.name, func(t *testing.T) {
			// 测试基本包信息获取（应该在所有镜像都工作）
			pkg, err := mirror.registry.GetPackageInformation(ctx, "react")
			if err != nil {
				t.Logf("镜像 %s 获取包信息失败: %v", mirror.name, err)
				return
			}
			assert.NotNil(t, pkg)
			assert.Equal(t, "react", pkg.Name)
		})
	}
}

// 边界情况测试
func TestSearchPackagesEdgeCases(t *testing.T) {
	registry := NewRegistry()
	ctx := context.Background()

	// 测试空查询
	result, err := registry.SearchPackages(ctx, "", 5)
	// 空查询可能返回错误或空结果，这取决于NPM API的行为
	if err == nil {
		assert.NotNil(t, result)
	}

	// 测试负数限制
	result, err = registry.SearchPackages(ctx, "test", -1)
	if err == nil {
		assert.NotNil(t, result)
		// 负数限制应该被转换为默认值20
	}
}

func TestGetPackageVersionEdgeCases(t *testing.T) {
	registry := NewRegistry()
	ctx := context.Background()

	// 测试不存在的包
	_, err := registry.GetPackageVersion(ctx, "this-package-definitely-does-not-exist-12345", "1.0.0")
	assert.Error(t, err, "不存在的包应该返回错误")

	// 测试不存在的版本
	_, err = registry.GetPackageVersion(ctx, "react", "999.999.999")
	assert.Error(t, err, "不存在的版本应该返回错误")
}

func TestGetDownloadStatsEdgeCases(t *testing.T) {
	registry := NewRegistry()
	ctx := context.Background()

	// 测试不存在的包
	_, err := registry.GetDownloadStats(ctx, "this-package-definitely-does-not-exist-12345", "last-day")
	// 下载统计API可能对不存在的包返回0或错误
	if err == nil {
		// 如果没有错误，说明API返回了某种默认值
		t.Log("下载统计API对不存在的包返回了结果（可能是0）")
	}

	// 测试无效的时间周期
	_, err = registry.GetDownloadStats(ctx, "react", "invalid-period")
	assert.Error(t, err, "无效的时间周期应该返回错误")
}
