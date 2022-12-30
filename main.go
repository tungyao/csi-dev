package main

import "tungyao/csi-dev/driver"

// TODO this error
func main() {
	ser := driver.NewNonBlockingGRPCServer()
	ser.Start("tcp://0.0.0.0:8000", &driver.LIdentityServer{
		Name:    "hello-csi",
		Version: "hello.world.csi",
	}, nil, &driver.LNodeServer{
		Driver: &driver.Nfs{},
	}, false)
	ser.Wait()
}
