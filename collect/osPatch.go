package collect

import (
	"github.com/StackExchange/wmi"
	"systemCollectClient/handle"
)

var Patch []handle.Patches

func GetPatchInfo() []handle.Patches {
	_ = wmi.Query("SELECT * FROM Win32_QuickFixEngineering", &Patch)
	return Patch
}
