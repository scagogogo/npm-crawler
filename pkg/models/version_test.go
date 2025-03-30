package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	// 创建一个测试版本
	version := &Version{
		Name:        "test-package",
		Version:     "1.0.0",
		Description: "Test package version 1.0.0",
		Main:        "index.js",
		Homepage:    "https://example.com/test-package",
		License:     "MIT",
		Repository: &Repository{
			Type: "git",
			URL:  "git+https://github.com/example/test-package.git",
		},
		Dist: &Dist{
			Shasum:  "abc123",
			Tarball: "https://registry.npmjs.org/test-package/-/test-package-1.0.0.tgz",
		},
		Author: &User{
			Name:  "Test Author",
			Email: "test@example.com",
		},
		Bugs: &Bugs{
			URL: "https://github.com/example/test-package/issues",
		},
	}

	// 基本属性测试
	assert.Equal(t, "test-package", version.Name)
	assert.Equal(t, "1.0.0", version.Version)
	assert.Equal(t, "Test package version 1.0.0", version.Description)
	assert.Equal(t, "index.js", version.Main)
	assert.Equal(t, "MIT", version.License)

	// 测试嵌套结构
	assert.NotNil(t, version.Repository)
	assert.Equal(t, "git", version.Repository.Type)
	assert.Equal(t, "git+https://github.com/example/test-package.git", version.Repository.URL)

	assert.NotNil(t, version.Dist)
	assert.Equal(t, "abc123", version.Dist.Shasum)
	assert.Equal(t, "https://registry.npmjs.org/test-package/-/test-package-1.0.0.tgz", version.Dist.Tarball)

	assert.NotNil(t, version.Author)
	assert.Equal(t, "Test Author", version.Author.Name)
	assert.Equal(t, "test@example.com", version.Author.Email)

	assert.NotNil(t, version.Bugs)
	assert.Equal(t, "https://github.com/example/test-package/issues", version.Bugs.URL)
}

func TestVersionJsonMarshaling(t *testing.T) {
	// 创建一个测试版本
	version := &Version{
		Name:        "test-package",
		Version:     "1.0.0",
		Description: "Test package version 1.0.0",
		Dependencies: map[string]string{
			"react":      "^17.0.0",
			"react-dom":  "^17.0.0",
			"typescript": "^4.0.0",
		},
		DevDependencies: map[string]string{
			"jest":    "^27.0.0",
			"webpack": "^5.0.0",
			"eslint":  "^8.0.0",
		},
		Dist: &Dist{
			Shasum:  "abc123",
			Tarball: "https://registry.npmjs.org/test-package/-/test-package-1.0.0.tgz",
		},
	}

	// 转换为 JSON
	bytes, err := json.Marshal(version)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// 从 JSON 解析回对象
	var parsedVersion Version
	err = json.Unmarshal(bytes, &parsedVersion)
	assert.Nil(t, err)

	// 验证字段
	assert.Equal(t, version.Name, parsedVersion.Name)
	assert.Equal(t, version.Version, parsedVersion.Version)
	assert.Equal(t, version.Description, parsedVersion.Description)

	// 验证依赖映射
	assert.Equal(t, 3, len(parsedVersion.Dependencies))
	assert.Equal(t, "^17.0.0", parsedVersion.Dependencies["react"])
	assert.Equal(t, "^17.0.0", parsedVersion.Dependencies["react-dom"])
	assert.Equal(t, "^4.0.0", parsedVersion.Dependencies["typescript"])

	assert.Equal(t, 3, len(parsedVersion.DevDependencies))
	assert.Equal(t, "^27.0.0", parsedVersion.DevDependencies["jest"])
	assert.Equal(t, "^5.0.0", parsedVersion.DevDependencies["webpack"])
	assert.Equal(t, "^8.0.0", parsedVersion.DevDependencies["eslint"])

	// 验证 Dist 对象
	assert.NotNil(t, parsedVersion.Dist)
	assert.Equal(t, version.Dist.Shasum, parsedVersion.Dist.Shasum)
	assert.Equal(t, version.Dist.Tarball, parsedVersion.Dist.Tarball)
}

func TestVersionFromJson(t *testing.T) {
	// 测试从 JSON 字符串解析
	jsonStr := `{
		"name": "lodash",
		"version": "4.17.21",
		"description": "Lodash modular utilities.",
		"main": "lodash.js",
		"author": {
			"name": "John-David Dalton",
			"email": "john.david.dalton@gmail.com"
		},
		"repository": {
			"type": "git",
			"url": "git+https://github.com/lodash/lodash.git"
		},
		"license": "MIT",
		"dependencies": {},
		"devDependencies": {
			"eslint": "^7.0.0",
			"mocha": "^8.0.0"
		},
		"dist": {
			"shasum": "79c399428f79c93e50e9f2942e0d50c7763edfc7",
			"tarball": "https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz"
		}
	}`

	var version Version
	err := json.Unmarshal([]byte(jsonStr), &version)
	assert.Nil(t, err)

	// 验证字段
	assert.Equal(t, "lodash", version.Name)
	assert.Equal(t, "4.17.21", version.Version)
	assert.Equal(t, "Lodash modular utilities.", version.Description)
	assert.Equal(t, "lodash.js", version.Main)
	assert.Equal(t, "MIT", version.License)

	// 验证嵌套结构
	assert.NotNil(t, version.Author)
	assert.Equal(t, "John-David Dalton", version.Author.Name)
	assert.Equal(t, "john.david.dalton@gmail.com", version.Author.Email)

	assert.NotNil(t, version.Repository)
	assert.Equal(t, "git", version.Repository.Type)
	assert.Equal(t, "git+https://github.com/lodash/lodash.git", version.Repository.URL)

	// 验证 DevDependencies
	assert.Equal(t, 2, len(version.DevDependencies))
	assert.Equal(t, "^7.0.0", version.DevDependencies["eslint"])
	assert.Equal(t, "^8.0.0", version.DevDependencies["mocha"])

	// 验证 Dist 对象
	assert.NotNil(t, version.Dist)
	assert.Equal(t, "79c399428f79c93e50e9f2942e0d50c7763edfc7", version.Dist.Shasum)
	assert.Equal(t, "https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz", version.Dist.Tarball)
}
