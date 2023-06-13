package driver

import (
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"sync"
)

// nfs operations

type Nfs struct {
	Addr string
	sync.Mutex
	FirstPath string // pre-mount path
}

// Mount the remote path on the local, in fact,
// and then create a directory on the local.
//The local directory will also sync to the remote one.

func (nfs *Nfs) mount(newPath string) error {
	nfs.Lock()
	defer nfs.Unlock()
	// newerPath is local path
	newerPath := "/home/dong/nfs/" + GetMd516([]byte(newPath))
	_, err := os.Stat(newerPath)
	isExist := true
	klog.Info(err)
	if err != nil {
		if os.IsNotExist(err) {
			isExist = false
		}
	}
	klog.Info("Open dir ", err)

	// create directory in the local
	if isExist == false {
		err = os.Mkdir(newerPath, 0777)
		klog.Info(err)
		if err != nil {
			return err
		}
		klog.Info("create dir ", newerPath)
	}
	klog.Info("nfs remote", nfs.Addr, nfs.FirstPath)
	// Mount remote path on the local path
	out, err := exec.Command("mount", nfs.Addr+":"+nfs.FirstPath, newerPath).CombinedOutput()
	klog.Info(string(out), err)
	if err != nil {
		klog.Info(err)
		return err
	}
	// 在刚刚那个目录上创建文件夹
	_, err = os.Stat(newerPath + "/" + newPath)
	if err == nil {
		return nil
	} else if os.IsNotExist(err) {
		if err = os.Mkdir(newerPath+"/"+newPath, 0777); err != nil {
			klog.Info(err)
			return err
		}
	}
	return nil
}

// 需要卸载本地目录 因为目录已经在远程创建好了
func (nfs *Nfs) unmount(path string) {
	newerPath := "/home/dong/nfs/" + GetMd516([]byte(path))
	_, err := exec.Command("umount", newerPath).CombinedOutput()
	if err != nil {
		klog.Info(err)
	}
}
