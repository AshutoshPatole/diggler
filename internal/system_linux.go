//go:build linux
// +build linux

package internal

import (
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GetSELinuxInfo() {
	selinuxStatus, err := os.ReadFile("/sys/fs/selinux/enforce")
	if err != nil {
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("SELinux Information")
	t.AppendHeader(table.Row{"Category", "Value"})
	t.AppendRows([]table.Row{
		{"SELinux Status", string(selinuxStatus)},
		{"Enforce Mode", filepath.Base(string(selinuxStatus))},
	})
	t.Render()
}

func GetAppArmorInfo() {
	profile, err := os.ReadFile("/proc/self/attr/current")
	if err != nil {
		return
	}
	if string(profile) != "unconfined" {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("AppArmor Information")
		t.AppendHeader(table.Row{"Category", "Value"})
		t.AppendRows([]table.Row{
			{"AppArmor Status", "Enabled"},
		})
		t.Render()
	}
}
