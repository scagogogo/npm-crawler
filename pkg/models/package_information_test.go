package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	// 创建一个测试包
	pkg := &Package{
		ID:          "test-package",
		Rev:         "1-abc123",
		Name:        "test-package",
		Description: "A test package for unit testing",
		DistTags: &DistTags{
			Latest: "1.0.0",
			Next:   "1.1.0",
		},
		Versions: map[string]*Version{
			"1.0.0": {
				Name:        "test-package",
				Version:     "1.0.0",
				Description: "Test package v1.0.0",
				Dist: &Dist{
					Shasum:  "abc123",
					Tarball: "https://registry.npmjs.org/test-package/-/test-package-1.0.0.tgz",
				},
			},
		},
		Readme: "# Test Package\nThis is a test package for unit testing.",
		Maintainers: []*User{
			{
				Name:  "tester",
				Email: "tester@example.com",
			},
		},
		Time: map[string]time.Time{
			"created":  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			"modified": time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			"1.0.0":    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Homepage: "https://example.com/test-package",
		Keywords: []string{"test", "package"},
		License:  "MIT",
		Repository: &Repository{
			Type: "git",
			URL:  "git+https://github.com/example/test-package.git",
		},
		Author: &Author{
			Name: "Test Author",
		},
		Bugs: &Bugs{
			URL: "https://github.com/example/test-package/issues",
		},
		ReadmeFilename: "README.md",
		Users: map[string]bool{
			"user1": true,
			"user2": true,
		},
		Contributors: []*Contributor{
			{
				Name: "contributor1",
				URL:  "https://github.com/contributor1",
			},
		},
	}

	// 基本属性测试
	assert.Equal(t, "test-package", pkg.ID)
	assert.Equal(t, "test-package", pkg.Name)
	assert.Equal(t, "A test package for unit testing", pkg.Description)

	// 测试嵌套结构
	assert.NotNil(t, pkg.DistTags)
	assert.Equal(t, "1.0.0", pkg.DistTags.Latest)

	// 测试版本映射
	assert.Contains(t, pkg.Versions, "1.0.0")
	assert.Equal(t, "test-package", pkg.Versions["1.0.0"].Name)
	assert.Equal(t, "1.0.0", pkg.Versions["1.0.0"].Version)

	// 测试 maintainers 数组
	assert.Len(t, pkg.Maintainers, 1)
	assert.Equal(t, "tester", pkg.Maintainers[0].Name)

	// 测试时间映射
	assert.Contains(t, pkg.Time, "created")
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), pkg.Time["created"])
}

func TestPackageToJsonString(t *testing.T) {
	// 创建一个简单的测试包
	pkg := &Package{
		ID:          "simple-package",
		Name:        "simple-package",
		Description: "Simple package for testing JSON conversion",
		DistTags: &DistTags{
			Latest: "1.0.0",
		},
	}

	// 测试 ToJsonString 方法
	jsonString := pkg.ToJsonString()
	assert.NotEmpty(t, jsonString)

	// 验证 JSON 是否有效
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	assert.Nil(t, err)

	// 验证关键字段
	assert.Equal(t, "simple-package", result["_id"])
	assert.Equal(t, "simple-package", result["name"])
	assert.Equal(t, "Simple package for testing JSON conversion", result["description"])

	// 验证嵌套结构
	distTags, ok := result["dist-tags"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "1.0.0", distTags["latest"])
}

func TestComplexPackageJsonConversion(t *testing.T) {
	// 创建一个完整的包对象
	now := time.Now().UTC()
	pkg := &Package{
		ID:          "complex-package",
		Rev:         "1-xyz789",
		Name:        "complex-package",
		Description: "A complex package for testing JSON conversion",
		DistTags: &DistTags{
			Latest: "2.0.0",
			Next:   "2.1.0-next.1",
		},
		Versions: map[string]*Version{
			"1.0.0": {
				Name:    "complex-package",
				Version: "1.0.0",
				Dist: &Dist{
					Shasum:  "xyz789",
					Tarball: "https://registry.npmjs.org/complex-package/-/complex-package-1.0.0.tgz",
				},
			},
			"2.0.0": {
				Name:    "complex-package",
				Version: "2.0.0",
				Dist: &Dist{
					Shasum:  "abc123",
					Tarball: "https://registry.npmjs.org/complex-package/-/complex-package-2.0.0.tgz",
				},
			},
		},
		Time: map[string]time.Time{
			"created":  now.Add(-30 * 24 * time.Hour),
			"modified": now,
			"1.0.0":    now.Add(-30 * 24 * time.Hour),
			"2.0.0":    now,
		},
	}

	// 转换为 JSON 字符串
	jsonString := pkg.ToJsonString()

	// 解析回 Package 对象
	var parsedPkg Package
	err := json.Unmarshal([]byte(jsonString), &parsedPkg)
	assert.Nil(t, err)

	// 验证关键字段是否正确转换
	assert.Equal(t, pkg.ID, parsedPkg.ID)
	assert.Equal(t, pkg.Name, parsedPkg.Name)
	assert.Equal(t, pkg.Description, parsedPkg.Description)

	// 验证嵌套结构
	assert.NotNil(t, parsedPkg.DistTags)
	assert.Equal(t, pkg.DistTags.Latest, parsedPkg.DistTags.Latest)
	assert.Equal(t, pkg.DistTags.Next, parsedPkg.DistTags.Next)

	// 验证版本映射
	assert.Contains(t, parsedPkg.Versions, "1.0.0")
	assert.Contains(t, parsedPkg.Versions, "2.0.0")
	assert.Equal(t, pkg.Versions["1.0.0"].Dist.Shasum, parsedPkg.Versions["1.0.0"].Dist.Shasum)
	assert.Equal(t, pkg.Versions["2.0.0"].Dist.Shasum, parsedPkg.Versions["2.0.0"].Dist.Shasum)
}
