package models

import (
	"encoding/json"
)

// RegistryInformation 表示 NPM Registry 的状态和信息
//
// 此结构包含了 NPM Registry 的基本状态信息，包括数据库信息、
// 文档数量、存储空间等统计数据。通常通过 Registry 客户端的
// GetRegistryInformation 方法获取。
//
// 主要字段说明:
//   - DbName: Registry 数据库名称
//   - DocCount: 当前存储的文档(包)总数
//   - DocDelCount: 已删除的文档数量
//   - UpdateSeq: 数据库更新序列号
//   - CompactRunning: 数据库是否正在压缩中
type RegistryInformation struct {
	DbName            string `json:"db_name"` // 数据库名称
	Engine            string `json:"engine"`
	DocCount          int    `json:"doc_count"`           // 文档(包)总数
	DocDelCount       int    `json:"doc_del_count"`       // 已删除的文档数
	UpdateSeq         int    `json:"update_seq"`          // 更新序列号
	PurgeSeq          int    `json:"purge_seq"`           // 清除序列号
	CompactRunning    bool   `json:"compact_running"`     // 是否正在压缩
	DiskSize          int64  `json:"disk_size"`           // 磁盘占用大小（字节）
	DataSize          int64  `json:"data_size"`           // 数据大小（字节）
	InstanceStartTime string `json:"instance_start_time"` // 实例启动时间戳

	// 存储大小相关信息
	Sizes struct {
		File     int64 `json:"file"`     // 文件大小（字节）
		Active   int64 `json:"active"`   // 活跃数据大小（字节）
		External int64 `json:"external"` // 外部数据大小（字节）
	} `json:"sizes"`

	// 其他字段，用于存储未明确定义的属性
	Other map[string]interface{} `json:"-"`

	DiskFormatVersion  int    `json:"disk_format_version"`
	CommittedUpdateSeq int    `json:"committed_update_seq"`
	CompactedSeq       int    `json:"compacted_seq"`
	UUID               string `json:"uuid"`
}

// ToJsonString 将 RegistryInformation 对象转换为 JSON 字符串
//
// 此方法将 Registry 信息序列化为 JSON 格式的字符串，方便存储或传输。
// 如果序列化过程中发生错误，将返回包含错误信息的字符串。
//
// 返回值:
//   - string: JSON 格式的字符串表示
//
// 示例:
//
//	info, err := registry.GetRegistryInformation(ctx)
//	if err == nil {
//		jsonStr := info.ToJsonString()
//		fmt.Println("Registry 信息:", jsonStr)
//	}
func (x *RegistryInformation) ToJsonString() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}
