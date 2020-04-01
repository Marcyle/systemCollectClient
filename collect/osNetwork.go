package collect

import (
	"github.com/StackExchange/wmi"
	"systemCollectClient/handle"
)

var Networks []handle.Interface

func GetOsNetworkInfo() []handle.Interface {
	_ = wmi.Query("Select * from Win32_NetworkAdapterConfiguration where IPEnabled = true", &Networks)
	return Networks
}
