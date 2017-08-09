package metric

//
type Query struct {
	obj string
}

// 打开查询器
func (p *Query) Open() error {
	//打开文件指针
	// 与windows相比，Linux在open时需要知道具体打开哪个文件
	// windows的query是通用指针，不需要提前搜集信息
	return nil
}

// 关闭查询器
func (p *Query) Close() error {
	// 关闭文件指针
	return nil
}

// 查询器采样
// 大多数计数器需要2次采样才能计算出数据
func (p *Query) Collect() error {
	// 与windows相比，winquery collect后系统给缓存了数据
	// 而linux需要自己缓存数据
	return nil
}

// 添加指标对象到查询器
func (p *Query) AddMetric(m interface{}) error {
	return nil
}
