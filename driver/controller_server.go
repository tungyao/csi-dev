package driver

import (
	"context"
	"csi-dev/csi"
	"k8s.io/klog/v2"
	"time"
)

// 在什么情况下会发射内存泄露 一种是临时泄露 一般发生在过长 string或slice的切片中 因为共用一个内存空间
// 永久内存泄露 发生祖师 time.Ticker未释放
type LControllerServer struct {
	*Nfs
	csi.ControllerServer
	LocalStorageSpaceName string
}

// slice扩容 小于阈值2倍 大于阈值 则通过 cap + 3 * 阈值 /4
// map扩容 装载因子 count/2^B 6.5 溢出的bucket = 2^B是否过多
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

	// 对已经关闭的chan读写 会发生什么 都能读 不能写 CAS compare and swap
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
	// 说说select
	// 用来做channel的多路复用 同一个时间只会有一个分支触发 可以安全的读写一个channel 向已满的channel或读没有数据的channel
	// select底层是一个scase  c hchan elem是个缓冲区地址
	// interface的实现 itab表示结构类型 data是个指针
}

// DeleteVolume
// 数据库的隔离级别 读未提交 其他事务能够读取到还未提交的数据，会出现脏读 读已提交 其他事务能够读取到事务已经提交的数据 会出现不可重复读 可重复度 比前一个好的地方 当前事务在重复读取时 指挥读取到一致的数据 可能会幻读
// 最后时 可串行化 有锁的保持 不会出现上面的问题
// 名称解释 脏读 A事务读取到了B事务更新的数据 然后B回滚了 不可重复读 A事务多次读取同一数据 事务B在A事务 读取过程中 多次更改了数据 事务A会读取到了不同的结果
// 幻读 在事务执行后 发现新数据没有进行修改
// 事务时基于 redo log 和 undo log 实现的
// undo log 是用来进行回滚的 事务产生的每条语句会被记录
// redo log 记录对数据的修改在事务进行中
// binlog 记录表的修改和数据的修改 可以在从主机上进行同步
// 缓存击穿 缓存失效了 请求直接到了数据源 解决方式 使用互斥锁 随机退避方式失效后 休眠一段时间 再次查询 缓存是加上随机时间 避免同时失效 还有一种凡是 不让缓存失效
// 缓存雪崩 缓存挂掉了 所有请求都会到达DB 熔断机制 或者 主从 集群 方式
// 缓存穿透 缓存和db都查询不命中的情况 一般使用布隆过滤器 或者给标记  布隆过滤器-> 多次hash求余 放入一个大数组中 查询时就能知道存不存在了
func (cs *LControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Info("get DeleteVolume", req.GetVolumeId())
	time.Now().Format(time.RFC1123)
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
				// 事务这里如果需要询问得深的话 则需要高出更多的东西
				// 悲观锁和乐观锁 最大区别就是 一个是在事务中一直加锁 一个是在提交时加锁
				// 死锁就是相互占用同一个资源
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER,
					},
				},
			},
		},
		// 索引下推 作用能够减少回表次数 操作方式 最左原则 匹配出来后 查询器可能会推出符合条件列数
		// redis
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
