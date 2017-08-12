package main

import (
	"github.com/cloudimmunity/system"
	"log"
)

func main() {
	sysInfo := system.GetSystemInfo()
	log.Printf("System Info: %#v", sysInfo)
	archName := system.MachineToArchName(sysInfo.Machine)
	log.Printf("Architecture Name: %#v", archName)
}
