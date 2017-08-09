package metric

import (
	"strings"
)

// 指标结构体
// object  : 对象名称,例如disk/partion/mem/cpu/network等等
// instance: 实例名称, 实例是对象的实例化，例如一个主机上的
//           disk对象有sda/sdb/sdc等等实例
// counter : 计数器名称，计数器是一个对象的指标属性，所有同一
//           对象实例化的实例都具有相同的计数器
type Metric struct {
	object   string
	instance string
	counter  string
}

// 生成指标对象的全名
// linux下生成的全名类似于disk.sda.writebytes
func (p *Metric) FullPath() string {
	if strings.TrimSpace(p.instance) == "" {
		return p.object + "." + p.counter
	}
	return p.object + "." + p.instance + "." + p.counter
}

// 检查指标全名是否合法
// 检查指标的对象、实例和计数器是否存在
func (p *Metric) CheckVaild() error {
	return nil
}

// 获取指标的值
// 大多数指标需要采样2次才能计算出数据
func (p *Metric) GetValue() (value float64, err error) {
	return
}

// 从查询器中移除指标
func (p *Metric) Remove() error {
	return nil
}

// 查询对象的实例和指标
func getInstancsAndMetrics(obj string) (i, c []string, err error) {
	return
}
