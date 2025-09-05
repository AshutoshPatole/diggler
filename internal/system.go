package internal

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetHostInfo() {
	h, _ := host.Info()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Host Information")
	t.AppendHeader(table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"Operating System", h.OS},
		{"Platform", h.Platform},
		{"Platform Version", h.PlatformVersion},
		{"Kernel Architecture", h.KernelArch},
		{"Kernel Version", h.KernelVersion},
	})
	if h.VirtualizationSystem == "" {
		t.AppendRow(table.Row{"Virtualization System", "Not Available"})
	} else {
		t.AppendRow(table.Row{"Virtualization System", h.VirtualizationSystem})
	}
	t.Render()
}

func GetCPUInfo() {
	c, _ := cpu.Info()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("CPU Information")
	t.AppendHeader(table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"CPU Cores", runtime.NumCPU()},
		{"CPU Model", c[0].ModelName},
		{"CPU Vendor", c[0].VendorID},
	})
	t.Render()
}

func GetMemoryInfo() {
	m, _ := mem.VirtualMemory()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Memory Information")
	t.AppendHeader(table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"Total Memory", fmt.Sprintf("%.2f GB", float64(m.Total)/1024/1024/1024)},
		{"Available Memory", fmt.Sprintf("%.2f GB", float64(m.Available)/1024/1024/1024)},
		{"Used Memory", fmt.Sprintf("%.2f GB", float64(m.Used)/1024/1024/1024)},
		{"Used Percentage", fmt.Sprintf("%.2f %%", m.UsedPercent)},
		{"Swap Total", fmt.Sprintf("%.2f GB", float64(m.SwapTotal)/1024/1024/1024)},
		{"Swap Free", fmt.Sprintf("%.2f GB", float64(m.SwapFree)/1024/1024/1024)},
		{"Huge Pages Total", m.HugePagesTotal},
		{"Huge Pages Free", m.HugePagesFree},
		{"Huge Pages Surp", m.HugePagesSurp},
		{"Huge Pages Size", m.HugePageSize},
		{"Anon Huge Pages", m.AnonHugePages},
	})
	t.Render()
}
