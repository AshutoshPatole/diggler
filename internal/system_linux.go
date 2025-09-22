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

func getSELinuxInfo() {
	selinuxStatus, err := runCommand("sestatus")
	if err != nil {
		return
	}
	fmt.Printf("\n\nSELinux Information:\n%s\n", selinuxStatus)
}

func getAppArmorInfo() {
	profile, err := runCommand("apparmor_status")
	if err != nil {
		return
	}
	fmt.Printf("\n\nAppArmor Information:\n%s\n", profile)
}

func GetSecurityInfo() {
	if isCommandAvailable("sestatus") {
		getSELinuxInfo()
	}
	if isCommandAvailable("apparmor_status") {
		getAppArmorInfo()
	}
}
