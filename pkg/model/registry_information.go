package model

type RegistryInformation struct {
	DbName         string `json:"db_name"`
	Engine         string `json:"engine"`
	DocCount       int    `json:"doc_count"`
	DocDelCount    int    `json:"doc_del_count"`
	UpdateSeq      int    `json:"update_seq"`
	PurgeSeq       int    `json:"purge_seq"`
	CompactRunning bool   `json:"compact_running"`
	Sizes          struct {
		Active   int64 `json:"active"`
		External int64 `json:"external"`
		File     int64 `json:"file"`
	} `json:"sizes"`
	DiskSize int64 `json:"disk_size"`
	DataSize int64 `json:"data_size"`
	Other    struct {
		DataSize int64 `json:"data_size"`
	} `json:"other"`
	InstanceStartTime  string `json:"instance_start_time"`
	DiskFormatVersion  int    `json:"disk_format_version"`
	CommittedUpdateSeq int    `json:"committed_update_seq"`
	CompactedSeq       int    `json:"compacted_seq"`
	UUID               string `json:"uuid"`
}
