package collect

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func GetDefenderInfo() []string {
	_ = ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, _ := oleutil.CreateObject("WbemScripting.SWbemLocator")
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface(ole.IID_IDispatch)
	defer wmi.Release()

	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer", nil, "root/SecurityCenter2")
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	resultRaw, _ := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM AntiVirusProduct")
	result := resultRaw.ToIDispatch()
	defer result.Release()

	itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", 0)
	item := itemRaw.ToIDispatch()
	defer item.Release()

	name, _ := oleutil.GetProperty(item, "displayName")
	path, _ := oleutil.GetProperty(item, "pathToSignedProductExe")

	var Stdout []string
	s1 := name.ToString()
	s2 := path.ToString()
	Stdout = append(Stdout, s1)
	Stdout = append(Stdout, s2)
	return Stdout
}
