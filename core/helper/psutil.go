package helper

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemStat struct {
	Hostname        string
	OS              string
	CpuPercent      float64
	MemoryTotal     uint64
	MemoryAvailable uint64
	MemoryUsed      uint64
	MemoryPercent   float64
	DiskTotal       uint64
	DiskFree        uint64
	DiskUsedPercent float64
	NetBytesRecv    uint64
	NetBytesSent    uint64
}

func GetSystemStat() *SystemStat {

	cp, _ := cpu.Percent(time.Second, false)
	v, _ := mem.VirtualMemory()
	n, _ := net.IOCounters(true)
	d, _ := disk.Usage("/")
	h, _ := host.Info()

	return &SystemStat{
		Hostname:        h.Hostname,
		OS:              h.OS,
		CpuPercent:      cp[0],
		MemoryTotal:     v.Total,
		MemoryAvailable: v.Available,
		MemoryUsed:      v.Used,
		MemoryPercent:   v.UsedPercent,
		DiskTotal:       d.Total,
		DiskFree:        d.Free,
		DiskUsedPercent: d.UsedPercent,
		NetBytesRecv:    n[0].BytesRecv,
		NetBytesSent:    n[0].BytesSent,
	}

}

func PrintSystemStat() {

	c, _ := cpu.Info()
	cp, _ := cpu.Percent(time.Second, false)
	v, _ := mem.VirtualMemory()
	n, _ := net.IOCounters(true)
	d, _ := disk.Usage("/")
	h, _ := host.Info()

	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")

	for _, sub_cpu := range c {
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		fmt.Printf("CPU        : %v %v cores \n", modelname, cores)
	}

	fmt.Printf("CPU Used   : %f%% \n", cp[0])
	fmt.Printf("Mem        : %v MB, Free: %v MB, Used: %v MB, Usage: %f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	fmt.Printf("HD         : %v GB, Free: %v GB, Usage: %f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("Network    : %v bytes / %v bytes \n", n[0].BytesRecv, n[0].BytesSent)
	fmt.Printf("OS         : %v(%v) %v \n", h.Platform, h.PlatformFamily, h.PlatformVersion)
	fmt.Printf("Hostname   : %v  \n", h.Hostname)
	fmt.Printf("SystemBoot : %v \n", btime)

}
