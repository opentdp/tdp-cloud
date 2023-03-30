package psutil

import (
	"encoding/json"
)

type SummaryStat struct {
	CreateAt     int64
	HostId       string
	HostName     string
	Uptime       uint64
	OS           string
	Platform     string
	KernelArch   string
	CpuCore      int
	CpuCoreLogic int
	CpuPercent   []float64
	MemoryTotal  uint64
	MemoryUsed   uint64
	PublicIpv4   string
	PublicIpv6   string
}

type DetailStat struct {
	*SummaryStat
	CpuModel      []string
	NetInterface  []NetInterface
	NetBytesRecv  uint64
	NetBytesSent  uint64
	DiskPartition []DiskPartition
	DiskTotal     uint64
	DiskUsed      uint64
	SwapTotal     uint64
	SwapUsed      uint64
}

type DiskPartition struct {
	Device     string
	Mountpoint string
	Fstype     string
	Total      uint64
	Used       uint64
}

type NetInterface struct {
	Name      string
	BytesRecv uint64
	BytesSent uint64
	Dropin    uint64
	Dropout   uint64
	Ipv4List  []string
	Ipv6List  []string
}

func (p *SummaryStat) From(s string) {
	json.Unmarshal([]byte(s), p)
}

func (p *SummaryStat) String() string {
	jsonbyte, _ := json.Marshal(p)
	return string(jsonbyte)
}
