// winmetric project winmetric.go
package metric

import (
	"errors"
	"fmt"
	"metric/win"
	"strings"
)

// 指标结构体
// object  : 对象名称,例如PhyscialDisk/LogicalDisk/Processor等等
// instance: 实例名称, 实例是对象的实例化，例如一个主机上的
//           LogicalDisk对象有C:/D:/E:等等实例
// counter : 计数器名称，计数器是一个对象的指标属性，所有同一
//           对象实例化的实例都具有相同的计数器
type Metric struct {
	object   string
	instance string
	counter  string
	handle   win.PDH_HCOUNTER
	value    win.PDH_FMT_COUNTERVALUE_DOUBLE
	rawvalue win.PDH_RAW_COUNTER
}

// 生成指标对象的全路径
// Windows下生成的指标路径类似于\LogicalDisk(C:)\% Idle Time
func (p *Metric) FullPath() string {
	if strings.TrimSpace(p.instance) == "" {
		return "\\" + p.object + "\\" + p.counter
	}
	return "\\" + p.object + "(" + p.instance + ")" + "\\" + p.counter
}

// 检查指标对象的全路径是否合法
// 检查指标的对象、实例和计数器是否存在
func (p *Metric) CheckVaild() error {
	var (
		err error
	)
	if win.PdhValidatePath(p.FullPath()) != win.ERROR_SUCCESS {
		err = errors.New("invaild counter full path")
	}
	return err
}

// 获取指标的值
// 大多数指标需要采样2次才能计算出数据
func (p *Metric) GetValue() (value float64, err error) {
	var (
		status uint32
		dwType uint32 = 0
	)

	status = win.PdhGetFormattedCounterValueDouble(p.handle, &dwType, &p.value)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("get format counter data with double type failed. ErrorCode: %0#8X", status)
		return
	}
	return p.value.DoubleValue, err
}

// 获取指标的原始数据
// 采样一次就行
func (p *Metric) GetRawValue() (value int64, err error) {
	var (
		status uint32
		dwType uint32 = 0
	)

	status = win.PdhGetRawCounterValue(p.handle, &dwType, &p.rawvalue)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("get raw counter data failed. ErrorCode: %0#8X", status)
		return
	}
	if p.rawvalue.CStatus != 0 {
		err = fmt.Errorf("get raw counter data failed. Cstatus: %d", p.rawvalue.CStatus)
	}
	return p.rawvalue.FirstValue, err
}

// 从查询器中移除指标
func (p *Metric) Remove() error {
	var (
		status uint32
		err    error
	)
	status = win.PdhRemoveCounter(p.handle)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("remove couter from query failed. couter full path:%s, ErrorCode:%0#8X",
			p.FullPath(), status,
		)
	}
	return err
}

// 获取对象的实例和计数器
func getInstancsAndMetrics(obj string) (i, c []string, err error) {
	var (
		status                 uint32
		pcchCounterListLength  uint32   = 0
		mszCounterList         []uint16 = make([]uint16, 1)
		pcchInstanceListLength uint32   = 0
		mszInstanceList        []uint16 = make([]uint16, 1)
		dwDetailLevel          uint32   = win.PERF_DETAIL_WIZARD
		dwFlags                uint32   = 0
	)

	status = win.PdhEnumObjectItems(
		0,
		0,
		obj,
		&mszCounterList[0],
		&pcchCounterListLength,
		&mszInstanceList[0],
		&pcchInstanceListLength,
		dwDetailLevel,
		dwFlags,
	)
	if status != win.PDH_MORE_DATA {
		err = fmt.Errorf("can't get object counters, get buffer size failed. ErrorCode: %0#8X", status)
		return
	}

	if pcchCounterListLength == 0 {
		err = fmt.Errorf("object %s have not counter", obj)
		return
	} else if pcchCounterListLength > 1 {
		mszCounterList = make([]uint16, pcchCounterListLength)
	}
	if pcchInstanceListLength > 1 {
		mszInstanceList = make([]uint16, pcchInstanceListLength)
	}

	status = win.PdhEnumObjectItems(
		0,
		0,
		obj,
		&mszCounterList[0],
		&pcchCounterListLength,
		&mszInstanceList[0],
		&pcchInstanceListLength,
		dwDetailLevel,
		dwFlags,
	)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("can't get object counters, get buffer failed. ErrorCode: %0#8X.", status)
		return
	}

	// 过滤数组中元素为空格或则空字符串
	for _, v := range win.UTF16ToString(mszInstanceList) {
		if strings.TrimSpace(v) != "" {
			i = append(i, v)
		}
	}

	for _, v := range win.UTF16ToString(mszCounterList) {
		if strings.TrimSpace(v) != "" {
			c = append(c, v)
		}
	}
	return i, c, nil
}
