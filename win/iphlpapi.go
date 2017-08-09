// winmetric project winmetric.go

// +build windows
package win

import (
	"syscall"
	"unsafe"
)

const (
	GAA_FLAG_INCLUDE_PREFIX = 0x00000010
	ERROR_BUFFER_OVERFLOW   = 111
)

type SocketAddress struct {
	Sockaddr       *syscall.RawSockaddrAny
	SockaddrLength int32
}

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

type IpAdapterAnycastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterAnycastAddress
	Address SocketAddress
}

type IpAdapterMulticastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterMulticastAddress
	Address SocketAddress
}

type IpAdapterDnsServerAdapter struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterDnsServerAdapter
	Address  SocketAddress
}

type IpAdapterPrefix struct {
	Length       uint32
	Flags        uint32
	Next         *IpAdapterPrefix
	Address      SocketAddress
	PrefixLength uint32
}

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

//
var (
	libiphlpapiDll                 *syscall.DLL
	iphelpapi_GetAdaptersAddresses *syscall.Proc
)

//
func init() {
	libiphlpapiDll = syscall.MustLoadDLL("iphlpapi.dll")

	iphelpapi_GetAdaptersAddresses = libiphlpapiDll.MustFindProc("GetAdaptersAddresses")

}

//
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
