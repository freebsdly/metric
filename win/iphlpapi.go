// winmetric project winmetric.go

// +build windows
package win

import (
	"syscall"
	"unsafe"
)

//
const (
	GAA_FLAG_INCLUDE_PREFIX = 0x00000010
	ERROR_BUFFER_OVERFLOW   = 111
)

// 定义套接字地址结构体
type SocketAddress struct {
	Sockaddr       *syscall.RawSockaddrAny
	SockaddrLength int32
}

// 定义IP网络适配器单播地址结构体
type IpAdapterUnicastAddress struct {
	Length             uint32
	Flags              uint32
	Next               *IpAdapterUnicastAddress
	Address            SocketAddress
	PrefixOrigin       int32
	SuffixOrigin       int32
	DadState           int32
	ValidLifetime      uint32
	PreferredLifetime  uint32
	LeaseLifetime      uint32
	OnLinkPrefixLength uint8
}

//
type IpAdapterAnycastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterAnycastAddress
	Address SocketAddress
}

// 定义IP网络适配器多播地址结构体
type IpAdapterMulticastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterMulticastAddress
	Address SocketAddress
}

//
type IpAdapterDnsServerAdapter struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterDnsServerAdapter
	Address  SocketAddress
}

//
type IpAdapterPrefix struct {
	Length       uint32
	Flags        uint32
	Next         *IpAdapterPrefix
	Address      SocketAddress
	PrefixLength uint32
}

//
type IpAdapterAddresses struct {
	Length                uint32
	IfIndex               uint32
	Next                  *IpAdapterAddresses
	AdapterName           *byte
	FirstUnicastAddress   *IpAdapterUnicastAddress
	FirstAnycastAddress   *IpAdapterAnycastAddress
	FirstMulticastAddress *IpAdapterMulticastAddress
	FirstDnsServerAddress *IpAdapterDnsServerAdapter
	DnsSuffix             *uint16
	Description           *uint16
	FriendlyName          *uint16
	PhysicalAddress       [syscall.MAX_ADAPTER_ADDRESS_LENGTH]byte
	PhysicalAddressLength uint32
	Flags                 uint32
	Mtu                   uint32
	IfType                uint32
	OperStatus            uint32
	Ipv6IfIndex           uint32
	ZoneIndices           [16]uint32
	FirstPrefix           *IpAdapterPrefix
	/* more fields might be present here. */
}

// 定义返回值常量
const (
	NO_ERROR                  = 0
	ERROR_NOT_SUPPORTED       = 50
	ERROR_INSUFFICIENT_BUFFER = 122
)

// 定义枚举类型
type MIB_TCP_STATE int

// TCP连接状态定义
const (
	MIB_TCP_STATE_CLOSED MIB_TCP_STATE = 1 + iota
	MIB_TCP_STATE_LISTEN
	MIB_TCP_STATE_SYN_SENT
	MIB_TCP_STATE_SYN_RCVD
	MIB_TCP_STATE_ESTAB
	MIB_TCP_STATE_FIN_WAIT1
	MIB_TCP_STATE_FIN_WAIT2
	MIB_TCP_STATE_CLOSE_WAIT
	MIB_TCP_STATE_CLOSING
	MIB_TCP_STATE_LAST_ACK
	MIB_TCP_STATE_TIME_WAIT
	MIB_TCP_STATE_DELETE_TCB
)

// TCP连接结构体
type MIB_TCPROW struct {
	DwState      uint32 // Old field used DWORD type.
	DwLocalAddr  uint32
	DwLocalPort  uint32
	DwRemoteAddr uint32
	DwRemotePort uint32
}

// 定义tcp连接表结构体，C中GetTcpTable需要结构体
// 的Table成员为元素数量为1的静态数组，如果这里设置
// 为go的切片，则无法获取到值，且会造成异常退出，这里
// 直接设置为元素个数为65535的数组
type MIB_TCPTABLE struct {
	DwNumEntries uint32
	Table        [65535]MIB_TCPROW
}

// 定义调用方法
var (
	libiphlpapiDll                 *syscall.DLL
	iphelpapi_GetAdaptersAddresses *syscall.Proc
	iphelpapi_GetTcpTable          *syscall.Proc
)

// 初始化调用方法
func init() {
	libiphlpapiDll = syscall.MustLoadDLL("iphlpapi.dll")

	iphelpapi_GetAdaptersAddresses = libiphlpapiDll.MustFindProc("GetAdaptersAddresses")
	iphelpapi_GetTcpTable = libiphlpapiDll.MustFindProc("GetTcpTable")

}

// 获取适配器地址
func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, iaa *IpAdapterAddresses, bufsize *uint32) uint32 {
	ret, _, _ := iphelpapi_GetAdaptersAddresses.Call(
		uintptr(family),
		uintptr(flags),
		reserved,
		uintptr(unsafe.Pointer(iaa)),
		uintptr(unsafe.Pointer(bufsize)),
	)

	return uint32(ret)
}

// 获取tcp连接表
func GetTcpTable(ptb *MIB_TCPTABLE, dwSize *uint32, border uint32) uint32 {
	ret, _, _ := iphelpapi_GetTcpTable.Call(
		uintptr(unsafe.Pointer(ptb)),
		uintptr(unsafe.Pointer(dwSize)),
		uintptr(border),
	)
	return uint32(ret)
}
