package main

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog/v2"
)

type Controller struct {
	csi.ControllerServer
}

func (Controller) CreateVolume(ctx context.Context, request *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.Infof("CreateVolume: called with args %#v", request)

	// 这里先返回一个假数据，模拟我们创建出了一块id为"qcow-1234567"容量为20G的云盘
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      "qcow-1234567",
			CapacityBytes: 0,
			VolumeContext: request.GetParameters(),
		},
	}, nil
}

func (Controller) DeleteVolume(ctx context.Context, request *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: called with args %#v", request)
	return &csi.DeleteVolumeResponse{}, nil
}

func (Controller) ControllerPublishVolume(ctx context.Context, request *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume: called with args %#v", request)
	pvInfo := map[string]string{DevicePathKey: "/dev/sdb"}
	return &csi.ControllerPublishVolumeResponse{PublishContext: pvInfo}, nil
}

func (Controller) ControllerUnpublishVolume(ctx context.Context, request *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.Infof("ControllerUnpublishVolume: called with args %#v", request)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (Controller) ValidateVolumeCapabilities(ctx context.Context, request *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	klog.Infof("ValidateVolumeCapabilities: called with args %#v", request)
	return &csi.ValidateVolumeCapabilitiesResponse{}, nil
}

func (Controller) ListVolumes(ctx context.Context, request *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	klog.Infof("ListVolumes: called with args %#v", request)
	return &csi.ListVolumesResponse{}, nil
}

func (Controller) GetCapacity(ctx context.Context, request *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	klog.Infof("GetCapacity: called with args %#v", request)
	return &csi.GetCapacityResponse{}, nil
}

var (
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	}
)

func (Controller) ControllerGetCapabilities(ctx context.Context, request *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	klog.Infof("ControllerGetCapabilities: called with args %#v", request)
	var caps []*csi.ControllerServiceCapability
	for _, controllerCap := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: controllerCap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (Controller) CreateSnapshot(ctx context.Context, request *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	klog.Infof("CreateSnapshot: called with args %#v", request)
	return &csi.CreateSnapshotResponse{}, nil
}

func (Controller) DeleteSnapshot(ctx context.Context, request *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	klog.Infof("DeleteSnapshot: called with args %#v", request)
	return &csi.DeleteSnapshotResponse{}, nil
}

func (Controller) ListSnapshots(ctx context.Context, request *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	klog.Infof("ListSnapshots: called with args %#v", request)
	return &csi.ListSnapshotsResponse{}, nil
}

func (Controller) ControllerExpandVolume(ctx context.Context, request *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume: called with args %#v", request)
	return &csi.ControllerExpandVolumeResponse{}, nil
}

func (Controller) ControllerGetVolume(ctx context.Context, request *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	klog.Infof("ControllerGetVolume: called with args %#v", request)
	return &csi.ControllerGetVolumeResponse{}, nil
}
