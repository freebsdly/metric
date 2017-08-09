// winmetric project winmetric.go

// +build windows
package win

import (
	"syscall"
	"unsafe"
)

//
var (
	psapiDLL                       *syscall.LazyDLL
	psapi_GetProcessImageFileNameW *syscall.LazyProc
)

//
func init() {
	psapiDLL = syscall.NewLazyDLL("psapi.dll")
	psapi_GetProcessImageFileNameW = psapiDLL.NewProc("GetProcessImageFileNameW")
}

//
func GetProcessImageFileName(hProcess syscall.Handle, lpImageFileName *uint16, nSize uint32) uint32 {
	ret, _, _ := psapi_GetProcessImageFileNameW.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpImageFileName)),
		uintptr(nSize),
	)

	return uint32(ret)
}
