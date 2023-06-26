package main

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog/v2"
)

type IdentityServer struct {
	csi.IdentityServer
}

func (i IdentityServer) GetPluginInfo(ctx context.Context, in *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	klog.Infof("GetPluginInfo: called with args %#v", in)
	return &csi.GetPluginInfoResponse{
		Name:          driverName,
		VendorVersion: version,
	}, nil
}

func (i IdentityServer) GetPluginCapabilities(ctx context.Context, in *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.Infof("GetPluginCapabilities: called with args %#v", in)
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

func (i IdentityServer) Probe(ctx context.Context, in *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	klog.Infof("Probe: called with args %#v", in)
	return &csi.ProbeResponse{}, nil
}
