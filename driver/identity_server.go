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
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
					},
				},
			},
		},
	}, nil
}

func (ids *LIdentityServer) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	klog.V(4).Infof("Probe: called with args %+v", req)
	return &csi.ProbeResponse{Ready: &wrappers.BoolValue{Value: ids.Status}}, nil
}
