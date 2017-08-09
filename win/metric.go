// winmetric project winmetric.go

// +build windows
package win

//定义windows性能对象
const (
	METRIC_SYSTEM                     string = "System"
	METRIC_PROCESSOR                  string = "Processor"
	METRIC_PROCESSOR_INFORMATION      string = "Processor Information"
	METRIC_PROCESSOR_PERFORMANCE      string = "ProcessorPerformance"
	METRIC_PHYSICALDISK               string = "PhysicalDisk"
	METRIC_LOGICALDISK                string = "LogicalDisk"
	METRIC_MEMORY                     string = "Memory"
	METRIC_NET_INTERFACE              string = "Network Interface"
	METRIC_NET_ADAPTER                string = "NETWork Adapter"
	METRIC_NET_QOS_POLICY             string = "Network QoS Policy"
	METRIC_NET_TCPV4                  string = "TCPv4"
	METRIC_NET_UDPV4                  string = "UDPv4"
	METRIC_PROCESS                    string = "Process"
	METRIC_DOTNET_CLR_DATA            string = ".NET CLR Data"
	METRIC_DOTNET_CLR_EXCEPTIONS      string = ".NET CLR Exceptions"
	METRIC_DOTNET_CLR_INTEROP         string = ".NET CLR Interop"
	METRIC_DOTNET_CLR_JIT             string = ".NET CLR Jit"
	METRIC_DOTNET_CLR_LOADING         string = ".NET CLR Loading"
	METRIC_DOTNET_CLR_LOCKSANDTHREADS string = ".NET CLR LocksAndThreads"
	METRIC_DOTNET_CLR_MEMORY          string = ".NET CLR Memory"
	METRIC_DOTNET_CLR_NETWORKING      string = ".NET CLR Networking"
	METRIC_DOTNET_CLR_REMOTING        string = ".NET CLR Remoting"
	METRIC_DOTNET_CLR_SECURITY        string = ".NET CLR Security"
)

// 定义windows常用实例名
const (
	INSTANCE_TOTAL string = "_Total"
)
