package collect

import (
	"github.com/StackExchange/wmi"
	"systemCollectClient/handle"
)

var SystemInfo []handle.OperaSystem

func GetOsVersionInfo() []handle.OperaSystem {
	_ = wmi.Query("Select * from Win32_OperatingSystem", &SystemInfo)
	return SystemInfo
}
