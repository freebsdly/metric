// winmetric project winmetric.go

// +build windows
package win

import (
	"syscall"
	"unsafe"
)

// error codes
const (
	ERROR_SUCCESS           = 0
	ERROR_INVALID_FUNCTION  = 1
	ERROR_INVALID_PARAMETER = 87
	ERROR_MORE_DATA         = 234
)

//
var (
	kernel32DLL                      *syscall.LazyDLL
	kernel32_GetLogicalDriveStringsW *syscall.LazyProc
	kernel32_QueryDosDeviceW         *syscall.LazyProc
)

//
func init() {
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	kernel32_GetLogicalDriveStringsW = kernel32DLL.NewProc("GetLogicalDriveStringsW")
	kernel32_QueryDosDeviceW = kernel32DLL.NewProc("QueryDosDeviceW")
}

// 获取逻辑磁盘名称
// 成功返回字符串实际长度
// 失败返回0
func GetLogicalDriveStrings(buflen uint32, buf *uint16) uint32 {
	ret, _, _ := kernel32_GetLogicalDriveStringsW.Call(
		uintptr(buflen),
		uintptr(unsafe.Pointer(buf)),
	)
	return uint32(ret)
}

// 查询dos设备名称
func QueryDosDevice(lpDeviceName *uint16, lpTargetPath *uint16, ucchMax uint32) uint32 {
	ret, _, _ := kernel32_QueryDosDeviceW.Call(
		uintptr(unsafe.Pointer(lpDeviceName)),
		uintptr(unsafe.Pointer(lpTargetPath)),
		uintptr(ucchMax),
	)
	return uint32(ret)

}
