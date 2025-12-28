package collector

type Configuration struct {
	CPUInfo    CPUInfo
	GPUInfo    GPUInfo
	RAMInfo    MemoryInfo
	DiskInfo   DiskPartition
	NetInfo    NetInfo
	TempInfo   TempInfo
	SystemInfo SystemInfo
}

type CPUInfo struct {
	ModelName   string   `json:"model_name"`
	Cores       int      `json:"cores"`   // physical
	Threads     int      `json:"threads"` // logical
	Percent     float64  `json:"percent"` // total percent
	Frequencies []int64  `json:"frequencies_mhz"`
	Flags       []string `json:"flags"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total_bytes"`
	Available   uint64  `json:"available_bytes"`
	Used        uint64  `json:"used_bytes"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskPartition struct {
	Device     string  `json:"device"`
	Mountpoint string  `json:"mountpoint"`
	Fstype     string  `json:"fstype"`
	Total      uint64  `json:"total_bytes"`
	Free       uint64  `json:"free_bytes"`
	Used       uint64  `json:"used_bytes"`
	UsedPct    float64 `json:"used_percent"`
}

type NetInfo struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
}

type GPUInfo struct {
	Name        string `json:"name"`
	Vendor      string `json:"vendor,omitempty"`
	Driver      string `json:"driver,omitempty"`
	MemoryTotal uint64 `json:"memory_total_bytes,omitempty"`
	MemoryUsed  uint64 `json:"memory_used_bytes,omitempty"`
	Utilization uint   `json:"utilization_percent,omitempty"`
	Temperature int    `json:"temperature_celsius,omitempty"`
}

type TempInfo struct {
	Sensor string  `json:"sensor"`
	Value  float64 `json:"value_celsius"`
}

type SystemInfo struct {
	Hostname    string `json:"hostname"`
	OS          string `json:"os"`
	Platform    string `json:"platform"`
	PlatformVer string `json:"platform_version"`
	UptimeSec   uint64 `json:"uptime_seconds"`
}
