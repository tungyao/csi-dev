package driver

import (
	"context"
	"csi-dev/csi"
	"github.com/golang/protobuf/ptypes/wrappers"
	"k8s.io/klog/v2"
)

// This file defines the plugin and returns its information.

// Getting identification

type LIdentityServer struct {
	csi.IdentityServer
	Name    string
	Version string
	Status  bool
}

func (ids *LIdentityServer) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	klog.Infof("get GetPluginInfo %#v", req)
	return &csi.GetPluginInfoResponse{
		Name:          ids.Name,
		VendorVersion: ids.Version,
	}, nil
}

func (ids *LIdentityServer) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.Infof("get GetPluginCapabilities %#v", req)
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

func (ids *LIdentityServer) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	klog.V(4).Infof("Probe: called with args %#v", req)
	return &csi.ProbeResponse{Ready: &wrappers.BoolValue{Value: ids.Status}}, nil
}
