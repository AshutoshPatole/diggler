//go:build windows

package internal

import (
	"os"
	"os/exec"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GetSecurityInfo() {
	// Not implemented on Windows
}

func FirewallStat() {
	// Not implemented on Windows
}

func GetDNSInfo() {
	t := NewTable("DNS Information", nil)
	dnsinfo, err := exec.Command("netsh", "interface", "ip", "show", "dns").Output()
	if err != nil {
		return
	}
	t.AppendRow(table.Row{string(dnsinfo)})
	t.Render()
}

func GetHostsFileInfo() {
	t := NewTable("Hosts File Information", nil)
	info, err := os.ReadFile("C:\\Windows\\System32\\drivers\\etc\\hosts")
	if err != nil {
		return
	}
	t.AppendRow(table.Row{string(info)})
	t.Render()
}
