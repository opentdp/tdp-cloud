package psutil

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemStat struct {
	HostId       string
	HostName     string
	Uptime       uint64
	OS           string
	CpuCore      int
	CpuPercent   float64
	MemoryTotal  uint64
	MemoryUsed   uint64
	DiskTotal    uint64
	DiskUsed     uint64
	NetBytesRecv uint64
	NetBytesSent uint64
}

func (p *SystemStat) From(s string) {

	json.Unmarshal([]byte(s), p)

}

func (p *SystemStat) String() string {

	jsonbyte, _ := json.Marshal(p)
	return string(jsonbyte)

}

func GetSystemStat() *SystemStat {

	cc, _ := cpu.Counts(true)
	cp, _ := cpu.Percent(time.Second, false)
	mv, _ := mem.VirtualMemory()
	dp, _ := disk.Partitions(false)
	ni, _ := net.IOCounters(true)
	hi, _ := host.Info()

	diskTotal := uint64(0)
	diskUsed := uint64(0)
	for _, dpi := range dp {
		du, _ := disk.Usage(dpi.Mountpoint)
		diskTotal += du.Total
		diskUsed += du.Used
	}

	netBytesRecv := uint64(0)
	netBytesSent := uint64(0)
	for _, nio := range ni {
		netBytesRecv += nio.BytesRecv
		netBytesSent += nio.BytesSent
	}

	return &SystemStat{
		HostId:       hi.HostID,
		HostName:     hi.Hostname,
		Uptime:       hi.Uptime,
		OS:           hi.OS,
		CpuCore:      cc,
		CpuPercent:   cp[0],
		MemoryTotal:  mv.Total,
		MemoryUsed:   mv.Used,
		DiskTotal:    diskTotal,
		DiskUsed:     diskUsed,
		NetBytesRecv: netBytesRecv,
		NetBytesSent: netBytesSent,
	}

}
