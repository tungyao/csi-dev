package driver

import (
	"context"
	"csi-dev/csi"
	"k8s.io/klog/v2"
)

type LControllerServer struct {
	*Nfs
	csi.ControllerServer
	LocalStorageSpaceName string
}

// CreateVolume 创建挂载地址
func (cs *LControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.Info("get CreateVolume")
	cs.LocalStorageSpaceName = req.GetName()
	klog.Info(req.GetName())
	// 将nfs挂载到该主机上
	err := cs.Nfs.mount(cs.LocalStorageSpaceName)
	defer cs.Nfs.unmount(cs.LocalStorageSpaceName)
	if err != nil {
		klog.Info(err)
		return nil, err
	}

	// 创建完成后卸载
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes:      req.CapacityRange.LimitBytes,
			VolumeId:           cs.LocalStorageSpaceName,
			VolumeContext:      req.GetParameters(),
			ContentSource:      req.GetVolumeContentSource(),
			AccessibleTopology: nil,
		},
	}, nil
}
func (cs *LControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Info("get DeleteVolume", req.GetVolumeId())
	return &csi.DeleteVolumeResponse{}, nil
}

//func (cs *LControllerServer) ControllerPublishVolume(context.Context, *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ControllerPublishVolume not implemented")
//}

//func (cs *LControllerServer) ControllerUnpublishVolume(context.Context, *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ControllerUnpublishVolume not implemented")
//}
//func (cs *LControllerServer) ValidateVolumeCapabilities(context.Context, *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ValidateVolumeCapabilities not implemented")
//}
//func (cs *LControllerServer) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ListVolumes not implemented")
//}
//func (cs *LControllerServer) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetCapacity not implemented")
//}

func (cs *LControllerServer) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: []*csi.ControllerServiceCapability{
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
					},
				},
			},
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER,
					},
				},
			},
		},
	}, nil
}

//func (cs *LControllerServer) CreateSnapshot(context.Context, *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
//}
//func (cs *LControllerServer) DeleteSnapshot(context.Context, *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshot not implemented")
//}
//func (cs *LControllerServer) ListSnapshots(context.Context, *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
//}
//func (cs *LControllerServer) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ControllerExpandVolume not implemented")
//}
//func (cs *LControllerServer) ControllerGetVolume(context.Context, *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ControllerGetVolume not implemented")
//}
