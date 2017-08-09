package metric

// Queryer封装了Query类型
type Queryer interface {
	Open() error
	Close() error
	Collect() error
	AddMetric(m interface{}) error
}

// 打开查询器
func OpenQuery(q Queryer) error {
	return q.Open()
}

// 关闭查询器
func CloseQuery(q Queryer) error {
	return q.Close()
}

// 查询器采样
// 大多数计数器需要2次采样才能计算出数据
func CollectQuery(q Queryer) error {
	return q.Collect()
}

//
func AddMetricToQuery(q Queryer, m interface{}) error {
	return q.AddMetric(m)
}

func NewQuery() *Query {
	return new(Query)
}

func NewQueryer() Queryer {
	var q Queryer
	q = new(Query)
	return q
}
