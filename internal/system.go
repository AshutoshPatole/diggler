package internal

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetHostInfo() {
	h, _ := host.Info()

	fmt.Printf("Operating System : %s\n", h.OS)

	fmt.Printf("Platform : %s\n", h.Platform)

	fmt.Printf("Platform Version : %s\n", h.PlatformVersion)

	fmt.Printf("Kernel Architecture : %s\n", h.KernelArch)

	fmt.Printf("Kernel Version : %s\n", h.KernelVersion)
	if h.VirtualizationSystem == "" {
		fmt.Printf("Virtualization System : Not Available\n")
	} else {
		fmt.Printf("Virtualization System : %s\n", h.VirtualizationSystem)
	}
}

func GetCPUInfo() {
	c, _ := cpu.Info()
	fmt.Printf("CPU Cores : %d\n", runtime.NumCPU())
	fmt.Printf("CPU Model : %s\n", c[0].ModelName)
	fmt.Printf("CPU Vendor : %s\n", c[0].VendorID)
}

func GetMemoryInfo() {
	m, _ := mem.VirtualMemory()
	fmt.Printf("Total Memory : %.2f GB\n", float64(m.Total)/1024/1024/1024)
	fmt.Printf("Available Memory : %.2f GB\n", float64(m.Available)/1024/1024/1024)
	fmt.Printf("Used Memory : %.2f GB\n", float64(m.Used)/1024/1024/1024)
	fmt.Printf("Used Percentage : %.2f %%\n", m.UsedPercent)

	fmt.Printf("Swap Total : %.2f GB\n", float64(m.SwapTotal)/1024/1024/1024)
	fmt.Printf("Swap Free : %.2f GB\n", float64(m.SwapFree)/1024/1024/1024)

	fmt.Printf("Huge Pages Total : %d\n", m.HugePagesTotal)
	fmt.Printf("Huge Pages Free : %d\n", m.HugePagesFree)
	fmt.Printf("Huge Pages Surp : %d\n", m.HugePagesSurp)
	fmt.Printf("Huge Pages Size : %d\n", m.HugePageSize)
	fmt.Printf("Anon Huge Pages : %d\n", m.AnonHugePages)
}
