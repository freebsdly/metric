//
// +build windows

package win

import (
	"syscall"
	"unsafe"
)

//
type SC_HANDLE uintptr

//
const (
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
)

//
const (
	SC_MANAGER_CONNECT            = 0x0001
	SC_MANAGER_CREATE_SERVICE     = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  = 0x0004
	SC_MANAGER_LOCK               = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG = 0x0020
	SC_MANAGER_ALL_ACCESS         = STANDARD_RIGHTS_REQUIRED |
		SC_MANAGER_CONNECT |
		SC_MANAGER_CREATE_SERVICE |
		SC_MANAGER_ENUMERATE_SERVICE |
		SC_MANAGER_LOCK |
		SC_MANAGER_QUERY_LOCK_STATUS |
		SC_MANAGER_MODIFY_BOOT_CONFIG
)

//
const (
	SERVICE_QUERY_CONFIG         = 0x0001
	SERVICE_CHANGE_CONFIG        = 0x0002
	SERVICE_QUERY_STATUS         = 0x0004
	SERVICE_ENUMERATE_DEPENDENTS = 0x0008
	SERVICE_START                = 0x0010
	SERVICE_STOP                 = 0x0020
	SERVICE_PAUSE_CONTINUE       = 0x0040
	SERVICE_INTERROGATE          = 0x0080
	SERVICE_USER_DEFINED_CONTROL = 0x0100
	SERVICE_ALL_ACCESS           = STANDARD_RIGHTS_REQUIRED |
		SERVICE_QUERY_CONFIG |
		SERVICE_CHANGE_CONFIG |
		SERVICE_QUERY_STATUS |
		SERVICE_ENUMERATE_DEPENDENTS |
		SERVICE_START |
		SERVICE_STOP |
		SERVICE_PAUSE_CONTINUE |
		SERVICE_INTERROGATE |
		SERVICE_USER_DEFINED_CONTROL
)

// dwServiceType
const (
	SERVICE_ADAPTER             = 0x00000004
	SERVICE_FILE_SYSTEM_DRIVER  = 0x00000002
	SERVICE_KERNEL_DRIVER       = 0x00000001
	SERVICE_RECOGNIZER_DRIVER   = 0x00000008
	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_USER_OWN_PROCESS    = 0x00000050
	SERVICE_USER_SHARE_PROCESS  = 0x00000060
	SERVICE_INTERACTIVE_PROCESS = 0x00000100
)

// dwStartType
const (
	SERVICE_BOOT_START   = 0x00000000
	SERVICE_SYSTEM_START = 0x00000001
	SERVICE_AUTO_START   = 0x00000002
	SERVICE_DEMAND_START = 0x00000003
	SERVICE_DISABLED     = 0x00000004
)

// dwErrorControl
const (
	SERVICE_ERROR_IGNORE   = 0x00000000
	SERVICE_ERROR_NORMAL   = 0x00000001
	SERVICE_ERROR_SEVERE   = 0x00000002
	SERVICE_ERROR_CRITICAL = 0x00000003
)

//
var (
	advapi32DLL                 *syscall.LazyDLL
	advapi32_OpenSCManagerW     *syscall.LazyProc
	advapi32_CloseServiceHandle *syscall.LazyProc
	advapi32_OpenServiceW       *syscall.LazyProc
	advapi32_CreateServiceW     *syscall.LazyProc
)

//
func init() {
	advapi32DLL = syscall.NewLazyDLL("Advapi32.dll")
	advapi32_OpenSCManagerW = advapi32DLL.NewProc("OpenSCManagerW")
	advapi32_CloseServiceHandle = advapi32DLL.NewProc("CloseServiceHandle")
	advapi32_OpenServiceW = advapi32DLL.NewProc("OpenServiceW")
	advapi32_CreateServiceW = advapi32DLL.NewProc("CreateServiceW")
}

// c中OpenManager返回的是一个结构体指针，不知道怎么转换，直接在go中返回通用指针
func OpenSCManager(lpMachineName, lpDatabaseName uintptr, dwDesiredAccess uint32) SC_HANDLE {
	ret, _, _ := advapi32_OpenSCManagerW.Call(lpMachineName, lpDatabaseName, uintptr(dwDesiredAccess))
	return SC_HANDLE(ret)
}

//
func CloseServiceHandle(schandle SC_HANDLE) uint32 {
	ret, _, _ := advapi32_CloseServiceHandle.Call(uintptr(schandle))
	return uint32(ret)
}

// Open a windows Service
func OpenService(hSCManager SC_HANDLE, lpServiceName string, dwDesiredAccess uint32) SC_HANDLE {
	sn := syscall.StringToUTF16Ptr(lpServiceName)
	ret, _, _ := advapi32_OpenServiceW.Call(uintptr(hSCManager), uintptr(unsafe.Pointer(sn)), uintptr(dwDesiredAccess))
	return SC_HANDLE(ret)
}

// Create a windows Service
func CreateService(
	hSCManager SC_HANDLE,
	lpServiceName string,
	lpDisplayName string,
	dwDesiredAccess uint32,
	dwServiceType uint32,
	dwStartType uint32,
	dwErrorControl uint32,
	lpBinaryPathName string,
	lpLoadOrderGroup string,
	lpDependencies string,
	lpServiceStartName string,
	lpPassword string,
) SC_HANDLE {

	var (
		lplog uintptr
		lpd   uintptr
		lpssn uintptr
		lppw  uintptr
		lpdti uintptr
	)
	lpsn := syscall.StringToUTF16Ptr(lpServiceName)
	lpdn := syscall.StringToUTF16Ptr(lpDisplayName)
	lpbn := syscall.StringToUTF16Ptr(lpBinaryPathName)
	//lpdti := uintptr(unsafe.Pointer(lpdwTagId))

	if lpLoadOrderGroup != "" {
		lplog = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpLoadOrderGroup)))
	}
	if lpDependencies != "" {
		lpd = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDependencies)))
	}
	if lpServiceStartName != "" {
		lpssn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpServiceStartName)))
	}
	if lpPassword != "" {
		lppw = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpPassword)))
	}

	ret, _, _ := advapi32_CreateServiceW.Call(
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(lpsn)),
		uintptr(unsafe.Pointer(lpdn)),
		uintptr(dwDesiredAccess),
		uintptr(dwServiceType),
		uintptr(dwStartType),
		uintptr(dwErrorControl),
		uintptr(unsafe.Pointer(lpbn)),
		lplog,
		lpdti,
		lpd,
		lpssn,
		lppw,
	)

	return SC_HANDLE(ret)
}
