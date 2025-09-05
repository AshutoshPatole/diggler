// +build linux

package internal

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/moby/docker/profiles/apparmor"
	"github.com/opencontainers/selinux/go-selinux"
)

func GetSELinuxInfo() {
	if selinux.GetEnabled() {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("SELinux Information")
		t.AppendHeader(table.Row{"Category", "Value"})
		t.AppendRows([]table.Row{
			{"SELinux Status", "Enabled"},
			{"Enforce Mode", selinux.EnforceMode()},
		})
		t.Render()
	}
}

func GetAppArmorInfo() {
	if apparmor.IsLoaded("") {
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
