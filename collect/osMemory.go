package collect

import (
	"github.com/StackExchange/wmi"
	"systemCollectClient/handle"
)

var Mem []handle.Memory

func GetOsMemoryInfo() []handle.Memory {
	_ = wmi.Query("Select * from Win32_PhysicalMemory", &Mem)
	return Mem
}
