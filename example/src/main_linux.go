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

	resolver := system.CallNumberResolver(archName)
	log.Printf("syscall 0: %#v", resolver(0))
	log.Printf("syscall 2: %#v", resolver(2))

	if yes, err := system.DefaultKernelFeatures.IsCompiled("CONFIG_FANOTIFY"); (err == nil) && yes {
		log.Println("FANOTIFY is enabled!")
	}
}
