package driver

import (
	"context"
	"github.com/google/uuid"
	"tungyao/csi-dev/csi"
)

type LControllerServer struct {
	Nfs
	csi.ControllerServer
	LocalStorageSpaceName string
}

func (cs *LControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	uid := uuid.New().String()
	cs.LocalStorageSpaceName = uid
	cs.Nfs.mount()
	defer Nfs.unmount()
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes:      req.CapacityRange.LimitBytes,
			VolumeId:           uid,
			VolumeContext:      req.GetParameters(),
			ContentSource:      req.GetVolumeContentSource(),
			AccessibleTopology: nil,
		},
	}, nil
}
func (cs *LControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	space, err := cs.GetSpace(req.VolumeId)
	if err != nil {
		return nil, err
	}
	space.Free()
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
