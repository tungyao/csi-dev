package driver

import (
	"csi-dev/csi"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"net"
	"os"
	"sync"
)

// NonBlockingGRPCServer Defines Non blocking GRPC server interfaces
type NonBlockingGRPCServer interface {
	// Start services at the endpoint
	Start(endpoint string, ids csi.IdentityServer, cs csi.ControllerServer, ns csi.NodeServer, testMode bool)
	// Wait for the service to stop
	Wait()
	// Stop the service gracefully
	Stop()
	// ForceStop Stops the service forcefully
	ForceStop()
}

func NewNonBlockingGRPCServer() NonBlockingGRPCServer {
	s := &nonBlockingGRPCServer{}
	return s
}

// NonBlocking server
type nonBlockingGRPCServer struct {
	wg     sync.WaitGroup
	server *grpc.Server
}

func (s *nonBlockingGRPCServer) Start(endpoint string, ids csi.IdentityServer, cs csi.ControllerServer, ns csi.NodeServer, testMode bool) {

	s.wg.Add(1)
	go s.serve(endpoint, ids, cs, ns, testMode)
}

func (s *nonBlockingGRPCServer) Wait() {
	s.wg.Wait()
}

func (s *nonBlockingGRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *nonBlockingGRPCServer) ForceStop() {
	s.server.Stop()
}

// 这里最好使用 uds

func (s *nonBlockingGRPCServer) serve(endpoint string, ids csi.IdentityServer, cs csi.ControllerServer, ns csi.NodeServer, testMode bool) {

	proto, addr, err := ParseEndpoint(endpoint)
	if err != nil {
		klog.Fatal(err.Error())
	}
	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			klog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}
	listener, err := net.Listen(proto, addr)
	if err != nil {
		klog.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(LogGRPC),
	}
	server := grpc.NewServer(opts...)
	s.server = server
	csi.RegisterIdentityServer(s.server, ids)
	csi.RegisterControllerServer(s.server, cs)
	csi.RegisterNodeServer(s.server, ns)
	klog.Infof("Listening for connections on address: %#v", listener.Addr())
	//reflection.Register(server)
	err = s.server.Serve(listener)
	if err != nil {
		klog.Fatalf("Failed to serve grpc server: %v", err)
	}
}
