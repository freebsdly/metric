package metric

//
type Metricer interface {
	FullPath() string
	CheckVaild() error
	GetValue() (v float64, err error)
	GetRawValue() (v int64, err error)
	Remove() error
}

// 获取指标对象的全名
func FullPath(m Metricer) string {
	return m.FullPath()
}

// 检查指标是否可用
func CheckVaild(m Metricer) error {
	return m.CheckVaild()
}

// 获取指标的值
func GetValue(m Metricer) (v float64, err error) {
	return m.GetValue()
}

//
func GetRawValue(m Metricer) (v int64, err error) {
	return m.GetRawValue()
}

func NewMetric(obj, inst, counter string) *Metric {
	m := new(Metric)
	m.object = obj
	m.instance = inst
	m.counter = counter
	return m
}

// 创建指标
func NewMetricer(obj, inst, counter string) Metricer {
	var mer Metricer
	mer = NewMetric(obj, inst, counter)
	return mer
}

// 移除指标
func RemoveMetric(m Metricer) error {
	return m.Remove()
}

// 获取性能对象的实例和指标
func GetInstancsAndMetrics(obj string) (i, c []string, err error) {
	return getInstancsAndMetrics(obj)
}

// 创建多个指标
func NewMetricers(obj string, insts, couts []string) []Metricer {
	var m []Metricer = make([]Metricer, 0)

	if len(insts) == 0 {
		insts = make([]string, 1)
	}
	for _, instance := range insts {
		for _, counter := range couts {
			metricer := NewMetricer(obj, instance, counter)
			m = append(m, metricer)
		}
	}
	return m
}
