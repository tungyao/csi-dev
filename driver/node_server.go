package driver

import (
	"context"
	"csi-dev/csi"
)

// TODO 需要对这些东西进行验证

type LNodeServer struct {
	csi.NodeServer
	Driver *Nfs
	NodeId string
}

// NodePublishVolume 实际的挂载操作
func (lns *LNodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	//targetPath := req.GetTargetPath()
	//lns.Driver.mount()

	return &csi.NodePublishVolumeResponse{}, nil
}
func (lns *LNodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return &csi.NodeUnpublishVolumeResponse{}, nil
}
func (lns *LNodeServer) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{NodeId: lns.NodeId}, nil
}
func (lns *LNodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{Capabilities: nil}, nil
}
