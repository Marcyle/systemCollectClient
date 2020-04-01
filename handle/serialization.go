package handle

import (
	"encoding/json"
)

type Serial interface {
	Serialization()
}

type Data struct {
	OsVersion  []OperaSystem
	OsDefender Defender
	CPU        []Processor
	Mem        []Memory
	Network    []Interface
	OsPatch    []Patches
}

type Defender struct {
	Name string
	Path string
}

type Processor struct {
	Name          string
	NumberOfCores uint32
}

type Memory struct {
	DeviceLocator string
	Capacity      uint64
	Manufacturer  string
	Caption       string
}

type Interface struct {
	Description string
	IPAddress   []string
	//DHCPServer string
	//DefaultIPGateway[] string
	MACAddress string
}

type Patches struct {
	HotFixID    string
	InstalledOn string
}

type OperaSystem struct {
	Caption        string
	Version        string
	InstallDate    string
	LastBootUpTime string
}

func (d *Data) serialization() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func Serializations(d *Data) string {
	s := d.serialization()
	return s
}
