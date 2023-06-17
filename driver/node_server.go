package driver

import (
	"context"
	"csi-dev/csi"
	"k8s.io/klog/v2"
)

// TODO 需要对这些东西进行验证

type LNodeServer struct {
	csi.NodeServer
	Driver *Nfs
	NodeId string
}

// NodeStageVolume 格式化硬盘，Mount到全局目录
func (lns *LNodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.V(4).Infof("NodeStageVolume: called with args %#v", req)

	return &csi.NodeStageVolumeResponse{}, nil
}

func (lns *LNodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.V(4).Infof("NodeUnstageVolume: called with args %#v", req)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

// NodePublishVolume 实际的挂载操作
func (lns *LNodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	//targetPath := req.GetTargetPath()
	//lns.Driver.mount()
	klog.Infof("node publish volume %#v", req)
	return &csi.NodePublishVolumeResponse{}, nil
}
func (lns *LNodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return &csi.NodeUnpublishVolumeResponse{}, nil
}
func (lns *LNodeServer) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{NodeId: lns.NodeId}, nil
}
func (lns *LNodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.Infof("NodeGetCapabilities: called with args %#v", req)

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
