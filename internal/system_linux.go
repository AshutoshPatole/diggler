//go:build linux

package internal

import (
	"os/exec"

	"github.com/jedib0t/go-pretty/v6/table"
)

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func GetSecurityInfo() {
	t := NewTable("Security Information", table.Row{"Category", "Value"})
	if isCommandAvailable("sestatus") {
		selinuxStatus, err := runCommand("sestatus")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"SELinux Information", selinuxStatus})
	}
	if isCommandAvailable("apparmor_status") {
		profile, err := runCommand("apparmor_status")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"AppArmor Information", profile})
	}
	t.Render()
}

func FirewallStat() {
	if isCommandAvailable("firewall-cmd") {
		t := NewTable("Firewall Status", table.Row{"Category", "Value"})
		status, err := runCommand("firewall-cmd", "--state")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"Status", status})

		rules, err := runCommand("firewall-cmd", "--list-all")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"Rules", rules})
		t.Render()
	}

	if isCommandAvailable("ufw") {
		t := NewTable("UFW Status", table.Row{"Category", "Value"})
		status, err := runCommand("ufw", "status")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"Status", status})

		rules, err := runCommand("ufw", "status", "verbose")
		if err != nil {
			return
		}
		t.AppendRow(table.Row{"Rules", rules})
		t.Render()
	}
}

func GetDNSInfo() {
	t := NewTable("DNS Information", nil)
	resolve, err := runCommand("grep", "-v", "#", "/etc/resolv.conf")
	if err != nil {
		return
	}
	t.AppendRow(table.Row{resolve})
	t.Render()
}

func GetHostsFileInfo() {
	t := NewTable("Hosts File Information", nil)
	resolve, err := runCommand("cat", "/etc/hosts")
	if err != nil {
		return
	}
	t.AppendRow(table.Row{resolve})
	t.Render()
}
