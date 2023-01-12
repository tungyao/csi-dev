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

	nodeId = flag.String("nodeid", "", "")
	flag.Parse()
	ser := driver.NewNonBlockingGRPCServer()
	nfs := &driver.Nfs{
		FirstPath: "/home/dong/nfs/nk",
		Addr:      "192.168.7.78",
	}
	ser.Start("unix://home/dong/project/csi-dev/csi.sock", &driver.LIdentityServer{
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
