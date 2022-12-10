package driver

import (
	"context"
	"os"
	"tungyao/csi-dev/csi"
	"tungyao/csi-dev/storage"
)

type LNodeServer struct {
	storage.LocalStorage
	csi.NodeServer
	NodeId string
}

// NodePublishVolume 实际的挂载操作
func (lns *LNodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	err := os.Link(req.VolumeId, targetPath)
	if err != nil {
		return nil, err
	}
	return &csi.NodePublishVolumeResponse{}, nil
}
func (lns *LNodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	space, err := lns.GetSpace(req.VolumeId)
	if err != nil {
		return nil, err
	}
	space.Free()
	return &csi.NodeUnpublishVolumeResponse{}, err
}
func (lns *LNodeServer) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{NodeId: lns.NodeId}, nil
}
func (lns *LNodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.CreateVolumeResponse{Volume: lns.}, nil
}
