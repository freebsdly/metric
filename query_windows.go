package metric

import (
	"errors"
	"fmt"
	"metric/win"
)

// Query是pdh_hquery方法的封装
type Query struct {
	handle win.PDH_HQUERY
}

// 打开查询器
func (p *Query) Open() error {
	var (
		status uint32
		err    error
	)
	status = win.PdhOpenQuery(0, 0, &p.handle)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("Open a query failed. ErrorCode: %0#8X.", status)
	}
	return err
}

// 关闭查询器
func (p *Query) Close() error {
	var (
		status uint32
		err    error
	)
	status = win.PdhCloseQuery(p.handle)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("Open a query failed. ErrorCode: %0#8X.", status)
	}
	return err

}

// 查询器采样
// 大多数计数器需要2次采样才能计算出数据
func (p *Query) Collect() error {
	var (
		status uint32
		err    error
	)
	status = win.PdhCollectQueryData(p.handle)
	if status != win.ERROR_SUCCESS {
		err = fmt.Errorf("collect query data failed. ErrorCode: %0#8X", status)
	}

	return err
}

// 添加指标到查询器中
func (p *Query) AddMetric(m interface{}) error {
	var (
		status uint32
		err    error
	)
	me, ok := m.(*Metric)
	if !ok {
		err = errors.New("can't convert arg to Metric type.")

	} else {
		status = win.PdhAddCounter(p.handle, me.FullPath(), 0, &me.handle)
		if status != win.ERROR_SUCCESS {
			err = fmt.Errorf("add couter to query failed. couter full path:%s, ErrorCode:%0#8X",
				me.FullPath(), status,
			)
		}
	}
	return err
}
