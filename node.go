package main

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog/v2"
)

type NodeServer struct {
	csi.NodeServer
}

// NodeStageVolume intro: create path on host and mount storage volume at this path
func (n *NodeServer) NodeStageVolume(ctx context.Context, request *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.Infof("NodeStageVolume: called with args %#v", request)
	return &csi.NodeStageVolumeResponse{}, nil
}

func (n *NodeServer) NodeUnstageVolume(ctx context.Context, request *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.Infof("NodeUnstageVolume: called with args %#v", request)
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (n *NodeServer) NodePublishVolume(ctx context.Context, request *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("NodePublishVolume: called with args %#v", request)
	return &csi.NodePublishVolumeResponse{}, nil

}

func (n *NodeServer) NodeUnpublishVolume(ctx context.Context, request *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Infof("NodeUnpublishVolume: called with args %#v", request)
	return &csi.NodeUnpublishVolumeResponse{}, nil

}

func (n *NodeServer) NodeGetVolumeStats(ctx context.Context, request *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	klog.Infof("NodeGetVolumeStats: called with args %#v", request)
	return &csi.NodeGetVolumeStatsResponse{}, nil

}

func (n *NodeServer) NodeExpandVolume(ctx context.Context, request *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	klog.Infof("NodeExpandVolume: called with args %#v", request)
	return &csi.NodeExpandVolumeResponse{}, nil

}

func (n *NodeServer) NodeGetCapabilities(ctx context.Context, request *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.Infof("NodeGetCapabilities: called with args %#v", request)
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (n *NodeServer) NodeGetInfo(ctx context.Context, request *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	klog.Infof("NodeGetInfo: called with args %#v", request)
	return &csi.NodeGetInfoResponse{
		NodeId: "1",
	}, nil
}
