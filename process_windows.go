package metric

import (
	"fmt"
	"metric/win"
	"syscall"
	"unsafe"
)

func AllProcesses() (ps map[uint32]string, err error) {
	var (
		snapshot syscall.Handle
		fullpath string
	)
	snapshot, err = syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return
	}
	defer syscall.CloseHandle(snapshot)
	var procEntry syscall.ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	if err = syscall.Process32First(snapshot, &procEntry); err != nil {
		return
	}
	ps = make(map[uint32]string)
	for {
		fullpath, err = GetProcessFullPath(procEntry.ProcessID)
		ps[procEntry.ProcessID] = fullpath
		err = syscall.Process32Next(snapshot, &procEntry)
		if err != nil {
			return
		}
	}
	return
}

// 获取进程的imagefile名称，注意：
// 名称的开头是dos格式的磁盘名称，例如\Device\HarddiskVolume3
// 转换为c:这样的名称请调用DosName进行转换
func GetProcessFullPath(pid uint32) (fullpath string, err error) {
	var (
		hprocess    syscall.Handle
		szImagePath []uint16 = make([]uint16, syscall.MAX_PATH)
	)
	hprocess, err = syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		return
	}
	num := win.GetProcessImageFileName(hprocess, &szImagePath[0], syscall.MAX_PATH)
	if num == 0 {
		err = fmt.Errorf("call getProcessImageFileName failed")
		return
	}
	syscall.CloseHandle(hprocess)
	fullpath = syscall.UTF16ToString(szImagePath[0:])
	return
}

// 获取dos下逻辑磁盘名称与nt系统逻辑磁盘名称的映射
func DosName() (names map[string]string, err error) {
	var (
		begin   uint32
		ldbuf   []uint16
		ldlen   uint32
		ddlen   uint32
		ldarray []uint16
	)
	ldbuf = make([]uint16, 100)
	ldlen = win.GetLogicalDriveStrings(100, &ldbuf[0])
	if ldlen == 0 {
		err = fmt.Errorf("call QueryDosDevice failed. return %d", ldlen)
		return
	}

	names = make(map[string]string)
	for i := uint32(4); i <= ldlen; i += 4 {
		ldarray = ldbuf[begin : i-1]
		ldarray[2] = 0
		buf1 := make([]uint16, 1000)
		ddlen = win.QueryDosDevice(&ldbuf[begin : i-1][0], &buf1[0], 1000)
		if ddlen == 0 || ddlen == win.ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		names[syscall.UTF16ToString(buf1[0:ddlen])] = syscall.UTF16ToString(ldarray)
		begin = i
	}
	return
}
