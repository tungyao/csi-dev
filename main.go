package main

import (
	"csi-dev/driver"
)

// TODO this error
func main() {
	ser := driver.NewNonBlockingGRPCServer()
	nfs := &driver.Nfs{
		FirstPath: "nk",
		Addr:      "192.168.7.78",
	}
	ser.Start("unix://home/dong/project/csi-dev/csi.sock", &driver.LIdentityServer{
		Name:    "hello-csi",
		Version: "hello.world.csi",
	}, &driver.LControllerServer{
		Nfs:                   nfs,
		LocalStorageSpaceName: "",
	}, &driver.LNodeServer{
		Driver: nfs,
	}, false)
	ser.Wait()
}

// 正确使用存储调配是关键 数据大部分是离线存储然后交给spark
// 如果使用sql spark
