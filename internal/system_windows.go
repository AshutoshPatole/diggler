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
	t := table.NewWriter()
	t.SetStyle(TABLE_STYLE)
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("DNS Information")
	dnsinfo, err := exec.Command("netsh", "interface", "ip", "show", "dns").Output()
	if err != nil {
		return
	}
	t.AppendRow(table.Row{string(dnsinfo)})
	t.Render()
}
