package collect

import (
	"systemCollectClient/handle"
)

func Collect() string {
	Stream := new(handle.Data)
	Stream.OsPatch = GetPatchInfo()
	Stream.OsVersion = GetOsVersionInfo()
	Stream.OsDefender.Name = GetDefenderInfo()[0]
	Stream.OsDefender.Path = GetDefenderInfo()[1]
	Stream.Mem = GetOsMemoryInfo()
	Stream.CPU = GetOsCpuInfo()
	Stream.Network = GetOsNetworkInfo()

	ss := handle.Serializations(Stream)
	return ss
}
