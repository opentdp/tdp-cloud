package psutil

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func Summary(withAddr bool) *SummaryStat {

	hi, _ := host.Info()
	cl, _ := cpu.Counts(true)
	cc, _ := cpu.Counts(false)
	cp, _ := cpu.Percent(time.Second, false)
	mv, _ := mem.VirtualMemory()

	stat := &SummaryStat{
		CreateAt:     time.Now().Unix(),
		HostId:       hi.HostID,
		HostName:     hi.Hostname,
		Uptime:       hi.Uptime,
		OS:           hi.OS,
		Platform:     hi.Platform,
		KernelArch:   hi.KernelArch,
		CpuCore:      cc,
		CpuCoreLogic: cl,
		CpuPercent:   cp,
		MemoryTotal:  mv.Total,
		MemoryUsed:   mv.Used,
	}

	if withAddr {
		ipv4, ipv6 := PublicAddress(false)
		stat.PublicIpv4 = ipv4
		stat.PublicIpv6 = ipv6
	}

	return stat

}

func Detail(withAddr bool) *DetailStat {

	ci, _ := cpu.Info()
	ni, _ := net.IOCounters(true)
	dp, _ := disk.Partitions(false)
	sw, _ := mem.SwapMemory()

	// CPU 信息

	cpuModel := []string{}
	for _, info := range ci {
		cpuModel = append(cpuModel, info.ModelName)
	}

	// 网络信息

	netInterface := []NetInterface{}
	netBytesRecv := uint64(0)
	netBytesSent := uint64(0)
	for _, nio := range ni {
		if nio.BytesRecv > 0 || nio.BytesSent > 0 {
			ift := NetInterface{
				Name:      nio.Name,
				BytesRecv: nio.BytesRecv,
				BytesSent: nio.BytesSent,
				Dropin:    nio.Dropin,
				Dropout:   nio.Dropout,
			}
			if withAddr {
				ift.Ipv4List, ift.Ipv6List = InterfaceAddrs(nio.Name)
			}
			netInterface = append(netInterface, ift)
		}
		netBytesRecv += nio.BytesRecv
		netBytesSent += nio.BytesSent
	}

	// 硬盘信息

	diskPartition := []DiskPartition{}
	diskTotaled := ","
	diskTotal := uint64(0)
	diskUsed := uint64(0)
	for _, dpi := range dp {
		du, _ := disk.Usage(dpi.Mountpoint)
		if du.Total > 0 || du.Used > 0 {
			diskPartition = append(diskPartition, DiskPartition{
				dpi.Device,
				dpi.Mountpoint, dpi.Fstype,
				du.Total, du.Used,
			})
		}
		if !strings.Contains(diskTotaled, dpi.Device) {
			diskTotaled += dpi.Device + ","
			diskTotal += du.Total
			diskUsed += du.Used
		}
	}

	// 汇总信息

	return &DetailStat{
		Summary(withAddr),
		cpuModel,
		netInterface,
		netBytesRecv,
		netBytesSent,
		diskPartition,
		diskTotal,
		diskUsed,
		sw.Total,
		sw.Used,
	}

}
