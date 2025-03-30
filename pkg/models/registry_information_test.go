package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegistryInformation(t *testing.T) {
	// 创建一个测试 Registry 信息
	registry := &RegistryInformation{
		DbName:            "registry",
		DocCount:          10000,
		DocDelCount:       100,
		UpdateSeq:         20000,
		PurgeSeq:          0,
		CompactRunning:    false,
		DiskSize:          1000000,
		DataSize:          900000,
		InstanceStartTime: "1641028800000",
	}

	// 设置 Sizes 结构体字段
	registry.Sizes.File = 1000000
	registry.Sizes.Active = 900000
	registry.Sizes.External = 500000

	// 基本属性测试
	assert.Equal(t, "registry", registry.DbName)
	assert.Equal(t, 10000, registry.DocCount)
	assert.Equal(t, 100, registry.DocDelCount)
	assert.Equal(t, 20000, registry.UpdateSeq)
	assert.Equal(t, 0, registry.PurgeSeq)
	assert.Equal(t, false, registry.CompactRunning)
	assert.Equal(t, int64(1000000), registry.DiskSize)
	assert.Equal(t, int64(900000), registry.DataSize)
	assert.Equal(t, "1641028800000", registry.InstanceStartTime)

	// 测试嵌套结构
	assert.Equal(t, int64(1000000), registry.Sizes.File)
	assert.Equal(t, int64(900000), registry.Sizes.Active)
	assert.Equal(t, int64(500000), registry.Sizes.External)
}

func TestRegistryInformationJsonMarshaling(t *testing.T) {
	// 创建一个测试 Registry 信息
	registry := &RegistryInformation{
		DbName:            "test-registry",
		DocCount:          5000,
		DocDelCount:       50,
		UpdateSeq:         10000,
		PurgeSeq:          0,
		CompactRunning:    false,
		DiskSize:          500000,
		DataSize:          450000,
		InstanceStartTime: time.Now().Format(time.RFC3339),
	}

	// 设置 Sizes 结构体字段
	registry.Sizes.File = 500000
	registry.Sizes.Active = 450000
	registry.Sizes.External = 250000

	// 转换为 JSON
	bytes, err := json.Marshal(registry)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// 从 JSON 解析回对象
	var parsedRegistry RegistryInformation
	err = json.Unmarshal(bytes, &parsedRegistry)
	assert.Nil(t, err)

	// 验证字段
	assert.Equal(t, registry.DbName, parsedRegistry.DbName)
	assert.Equal(t, registry.DocCount, parsedRegistry.DocCount)
	assert.Equal(t, registry.DocDelCount, parsedRegistry.DocDelCount)
	assert.Equal(t, registry.UpdateSeq, parsedRegistry.UpdateSeq)
	assert.Equal(t, registry.PurgeSeq, parsedRegistry.PurgeSeq)
	assert.Equal(t, registry.CompactRunning, parsedRegistry.CompactRunning)
	assert.Equal(t, registry.DiskSize, parsedRegistry.DiskSize)
	assert.Equal(t, registry.DataSize, parsedRegistry.DataSize)
	assert.Equal(t, registry.InstanceStartTime, parsedRegistry.InstanceStartTime)

	// 验证 Sizes 结构体字段
	assert.Equal(t, registry.Sizes.File, parsedRegistry.Sizes.File)
	assert.Equal(t, registry.Sizes.Active, parsedRegistry.Sizes.Active)
	assert.Equal(t, registry.Sizes.External, parsedRegistry.Sizes.External)
}

func TestRegistryInformationFromJson(t *testing.T) {
	// 测试从 JSON 字符串解析
	jsonStr := `{
		"db_name": "registry",
		"doc_count": 1000,
		"doc_del_count": 10,
		"update_seq": 5000,
		"purge_seq": 0,
		"compact_running": false,
		"disk_size": 1000000,
		"data_size": 900000,
		"instance_start_time": "1641028800000",
		"sizes": {
			"file": 1000000,
			"active": 900000,
			"external": 500000
		}
	}`

	var registry RegistryInformation
	err := json.Unmarshal([]byte(jsonStr), &registry)
	assert.Nil(t, err)

	// 验证字段
	assert.Equal(t, "registry", registry.DbName)
	assert.Equal(t, 1000, registry.DocCount)
	assert.Equal(t, 10, registry.DocDelCount)
	assert.Equal(t, 5000, registry.UpdateSeq)
	assert.Equal(t, 0, registry.PurgeSeq)
	assert.Equal(t, false, registry.CompactRunning)
	assert.Equal(t, int64(1000000), registry.DiskSize)
	assert.Equal(t, int64(900000), registry.DataSize)
	assert.Equal(t, "1641028800000", registry.InstanceStartTime)
	assert.Equal(t, int64(1000000), registry.Sizes.File)
	assert.Equal(t, int64(900000), registry.Sizes.Active)
	assert.Equal(t, int64(500000), registry.Sizes.External)

	// 测试 ToJsonString 方法
	jsonString := registry.ToJsonString()
	assert.NotEmpty(t, jsonString)
}
