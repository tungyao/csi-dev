package driver

import (
	"context"
	"csi-dev/csi"
	"github.com/golang/protobuf/ptypes/wrappers"
	"k8s.io/klog/v2"
)

// 这整个文件都是定义和返回这个插件的信息

// LIdentityServer 获取认证信息
type LIdentityServer struct {
	csi.IdentityServer
	Name    string
	Version string
	Status  bool
}

func (ids *LIdentityServer) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	klog.Info("get GetPluginInfo")
	return &csi.GetPluginInfoResponse{
		Name:          ids.Name,
		VendorVersion: ids.Version,
	}, nil
}

func (ids *LIdentityServer) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.Info("get GetPluginCapabilities")
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
		},
	}, nil
}

func (ids *LIdentityServer) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{Ready: &wrappers.BoolValue{Value: ids.Status}}, nil
}
