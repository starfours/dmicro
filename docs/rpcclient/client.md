# RPC 客户端

直接使用`drpc`的`Endpoint`就能创建一个可用的rpc客户端。
那为什么还需要`rpc client`呢？

前面的文档已经阐述了如何创建一个原始的`drpc.Endpoint`,链接到服务端。
但是一个成熟的微服务框架的客户端不仅仅是一个单独的链接就可行的。

它需要包含`服务发现`,`链路追踪`,`限流`,`指标上报`等等服务治理的功能。这些组件与`drpc.Endpoint`如何有机的结合起来，就需要`rpc client`组件来实现了。

## 快速开始

```go
func main() {
	serviceName := "foo"
	cli := client.NewRpcClient(serviceName)
	var result int
	stat := cli.Call("/math/add",
		[]int{1, 2, 3, 4, 5},
		&result,
		message.WithSetMeta("author", "liuzhiming"),
	).Status()
	if !stat.OK() {
		logger.Fatalf("%v", stat)
	}
	logger.Printf("result: %d", result)
}

```

## 支持的配置方法

## 使用组件
