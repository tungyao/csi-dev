package main

import (
	"csi-dev/driver"
	"flag"
)

var (
	nodeId *string
)

// TODO this error
func main() {

	nodeId = flag.String("nodeid", "abc", "")
	flag.Parse()
	ser := driver.NewNonBlockingGRPCServer()
	nfs := &driver.Nfs{
		FirstPath: "/mnt/nfs_share",
		Addr:      "192.168.7.102",
	}
	// 让我看看 pod的创建流程
	// 配置文件通知到apiserver apiserver通知controller-manager创建一个资源 然后将资源存储到etcd中
	// 调度器开始预选将节点
	// kubulet收到创建信息后 开始创建 创建成功后 将pod的运行信息发回调度器 然后存储到etcd中
	ser.Start("unix:///csi/csi.sock", &driver.LIdentityServer{
		//ser.Start("tcp://0.0.0.0:9000", &driver.LIdentityServer{
		Name:    "hello-csi",
		Version: "hello.world.csi",
		Status:  true,
	}, &driver.LControllerServer{
		Nfs:                   nfs,
		LocalStorageSpaceName: "",
	}, &driver.LNodeServer{
		Driver: nfs,
		NodeId: *nodeId,
	}, false)
	ser.Wait()
}

// 正确使用存储调配是关键 数据大部分是离线存储然后交给spark
// 如果使用sql spark
