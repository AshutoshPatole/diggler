//go:build windows

package internal

import (
	"fmt"
	"os/exec"
)

func GetSecurityInfo() {
	// Not implemented on Windows
}

func FirewallStat() {
	// Not implemented on Windows
}

func GetDNSInfo() {
	fmt.Printf("\n\nDNS\n")
	dnsinfo, err := exec.Command("netsh", "interface", "ip", "show", "dns").Output()
	if err != nil {
		return
	}
	fmt.Printf("%s\n", dnsinfo)
}
