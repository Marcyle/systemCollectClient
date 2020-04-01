package collect

import (
	"github.com/StackExchange/wmi"
	"systemCollectClient/handle"
)

var Cpu []handle.Processor

func GetOsCpuInfo() []handle.Processor {
	_ = wmi.Query("Select * from Win32_Processor", &Cpu)
	return Cpu
}
