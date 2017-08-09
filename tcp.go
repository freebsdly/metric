package metric

//
type TcpRow struct {
	LocalAddr  string
	LocalPort  uint16
	RemoteAddr string
	RemotePort uint16
	State      int
}

//
type TcpTable struct {
	Table []TcpRow
}

// 获取当前被监听的tcp端口号
func (p *TcpTable) GetActivePorts() []uint16 {
	var ports []uint16
	if len(p.Table) == 0 {
		return ports
	}
	var tports = make(map[uint16]string)
	for _, v := range p.Table {
		tports[v.LocalPort] = " "
	}

	for k, _ := range tports {
		ports = append(ports, k)
	}

	return ports

}
