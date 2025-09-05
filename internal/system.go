package internal

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func GetSystemInfo() {
	h, _ := host.Info()

	fmt.Println("Operating System : " + h.OS)

	fmt.Println("Platform : " + h.Platform)

	fmt.Println("Platform Version : " + h.PlatformVersion)

	fmt.Println("Kernel Architecture : " + h.KernelArch)

	fmt.Println("Kernel Version : " + h.KernelVersion)
	if h.VirtualizationSystem == "" {
		fmt.Println("Virtualization System : Not Available")
	} else {
		fmt.Println("Virtualization System : " + h.VirtualizationSystem)
	}

	fmt.Println(h)

}
