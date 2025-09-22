package internal

import (
	"fmt"
	"runtime"

	"github.com/beevik/ntp"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

var TABLE_STYLE table.Style = table.StyleBold


func GetHostInfo() {
	h, _ := host.Info()
	t := NewTable("Host Information", table.Row{"Category", "Value"})
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
	t := NewTable("CPU Information", table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"CPU Cores", runtime.NumCPU()},
		{"CPU Model", c[0].ModelName},
		{"CPU Vendor", c[0].VendorID},
	})
	t.Render()
}

func GetMemoryInfo() {
	m, _ := mem.VirtualMemory()
	t := NewTable("Memory Information", table.Row{"Category", "Value"})
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

func GetNTPInfo() {
	response, err := ntp.Query("pool.ntp.org")
	if err != nil {
		return
	}
	t := NewTable("NTP Information", table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"Clock Offset", response.ClockOffset},
		{"Stratum", response.Stratum},
		{"Leap", response.Leap},
		{"Root Delay", response.RootDelay},
		{"Root Dispersion", response.RootDispersion},
		{"Reference ID", response.ReferenceID},
		{"Reference Time", response.ReferenceTime},
		{"Poll", response.Poll},
		{"Precision", response.Precision},
	})
	t.Render()
}

func GetProcessInfo() {

	s := StartSpinner("Gathering process information...")
	processes, _ := process.Processes()
	t := NewTable("Process Information", table.Row{"PID", "Process", "Path"})
	for _, p := range processes {
		name, _ := p.Name()
		openFiles, err := p.OpenFiles()
		if err != nil {
			continue
		}
		for _, file := range openFiles {
			t.AppendRow(table.Row{p.Pid, name, file.Path})
		}
	}
	s.Stop()
	t.Render()
}

func GetConnections() {
	s := StartSpinner("Gathering network connections...")
	connections, _ := net.Connections("all")
	t := NewTable("Network Connections", table.Row{"PID", "Local Address", "Remote Address", "Status"})
	for _, conn := range connections {
		t.AppendRow(table.Row{conn.Pid, fmt.Sprintf("%s:%d", conn.Laddr.IP, conn.Laddr.Port), fmt.Sprintf("%s:%d", conn.Raddr.IP, conn.Raddr.Port), conn.Status})
	}
	s.Stop()
	t.Render()
}
