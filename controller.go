package main

import (
	"context"
	"errors"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"io/fs"
	"k8s.io/klog/v2"
	"os"
)

type Controller struct {
	csi.ControllerServer
	nfs *Nfs
}

var Volume = make(map[string]string)

// create disk then return volume info

func (c *Controller) CreateVolume(ctx context.Context, request *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.Infof("CreateVolume: called with args %#v", request)

	// 远程创建一个目录 create directory on remote
	dt := c.nfs.provisionalPath(request.Name)
	if dt.err != nil { // 错误创建
		klog.Infoln("CreateVolume provisionalPath", dt.err)
		return nil, dt.err
	}
	dt.createPath("/" + request.Name)
	// umount nfs and delete local dir

	err := c.nfs.umount(dt.localPath)
	if err != nil {
		klog.Infoln("CreateVolume umount", err)
		return nil, err
	}
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      dt.name,
			CapacityBytes: 0,
			VolumeContext: request.GetParameters(),
		},
	}, nil
}

func (c *Controller) DeleteVolume(ctx context.Context, request *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: called with args %#v", request)
	dt := c.nfs.provisionalPath(request.VolumeId)
	if dt.err != nil { // 错误创建
		klog.Infoln("provisionalPath", dt.err)
		return nil, dt.err
	}
	// 挂载远程目录
	//c.nfs.mount("", dt.localPath)
	err := os.RemoveAll(dt.localPath + "/" + request.VolumeId)
	klog.Infoln("DeleteVolume Remove", err)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		klog.Infoln("DeleteVolume Remove", err)
		return nil, err
	}
	err = c.nfs.umount(dt.localPath)

	// 还需要额外移除node上临时创建的目录
	err = dt.removeProvisionalPath()
	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerPublishVolume 附加卷 在nfs可以不做任何操作
func (Controller) ControllerPublishVolume(ctx context.Context, request *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume: called with args %#v", request)

	return &csi.ControllerPublishVolumeResponse{}, nil
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
		csi.ControllerServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER,
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
