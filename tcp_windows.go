package metric

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"metric/win"
	"net"
)

// 获取tcp连接表
func GetTcpTable() (tt *TcpTable, err error) {
	var ptb *win.MIB_TCPTABLE = &win.MIB_TCPTABLE{}
	var size uint32

	if win.GetTcpTable(ptb, &size, 1) != win.ERROR_INSUFFICIENT_BUFFER {
		return tt, fmt.Errorf("call native GetTcpTable failed. return is not ERROR_INSUFFICIENT_BUFFER")
	}

	if win.GetTcpTable(ptb, &size, 1) != win.NO_ERROR {
		return tt, fmt.Errorf("call native GetTcpTable failed. return is not NO_ERROR")
	}

	var tr = TcpRow{}
	tt = &TcpTable{Table: make([]TcpRow, ptb.DwNumEntries)}
	for i := uint32(0); i < ptb.DwNumEntries; i++ {
		tr.RemoteAddr = addrDecode(ptb.Table[i].DwRemoteAddr).String()
		tr.LocalAddr = addrDecode(ptb.Table[i].DwLocalAddr).String()
		tr.RemotePort = portDecode(ptb.Table[i].DwRemotePort)
		tr.LocalPort = portDecode(ptb.Table[i].DwLocalPort)
		tt.Table[i] = tr
	}
	return tt, nil

}

// 将uint32格式的IP地址转换为nete.IP形式的IP地址
func addrDecode(addr uint32) net.IP {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, addr)
	return net.IP(buf.Bytes())
}

// 将uint32格式的端口转换为uint16格式的端口号
func portDecode(port uint32) uint16 {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, uint16(port))
	return binary.BigEndian.Uint16(buf.Bytes())
}
