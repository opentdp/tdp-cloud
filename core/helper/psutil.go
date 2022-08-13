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

func SystemInfo() {

	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	v, _ := mem.VirtualMemory()
	nv, _ := net.IOCounters(true)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")

	for _, sub_cpu := range c {
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		fmt.Printf("CPU        : %v %v cores \n", modelname, cores)
	}
	fmt.Printf("CPU Used   : %f%% \n", cc[0])
	fmt.Printf("Mem        : %v MB, Free: %v MB, Used: %v MB, Usage: %f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	fmt.Printf("Network    : %v bytes / %v bytes \n", nv[0].BytesRecv, nv[0].BytesSent)
	fmt.Printf("HD         : %v GB, Free: %v GB, Usage: %f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("OS         : %v(%v) %v \n", n.Platform, n.PlatformFamily, n.PlatformVersion)
	fmt.Printf("Hostname   : %v  \n", n.Hostname)
	fmt.Printf("SystemBoot : %v \n", btime)

}
