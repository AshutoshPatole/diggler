//go:build linux

package internal

import (
	"fmt"
	"os/exec"
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
	if isCommandAvailable("sestatus") {
		selinuxStatus, err := runCommand("sestatus")
		if err != nil {
			return
		}
		fmt.Printf("\n\nSELinux Information:\n%s\n", selinuxStatus)
	}
	if isCommandAvailable("apparmor_status") {
		profile, err := runCommand("apparmor_status")
		if err != nil {
			return
		}
		fmt.Printf("\n\nAppArmor Information:\n%s\n", profile)
	}
}

func FirewallStat() {
	if isCommandAvailable("firewall-cmd") {
		fmt.Printf("\n\nFirewall Status")
		status, err := runCommand("firewall-cmd", "--state")
		if err != nil {
			return
		}
		fmt.Printf("Status: %s\n", status)

		rules, err := runCommand("firewall-cmd", "--list-all")
		if err != nil {
			return
		}
		fmt.Printf("Rules: %s\n", rules)
	}

	if isCommandAvailable("ufw") {

		fmt.Printf("\n\nUFW Status")
		status, err := runCommand("ufw", "status")
		if err != nil {
			return
		}
		fmt.Printf("Status: %s\n", status)

		rules, err := runCommand("ufw", "status", "verbose")
		if err != nil {
			return
		}
		fmt.Printf("Rules: %s\n", rules)
	}
}
