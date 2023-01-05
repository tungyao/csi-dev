package driver

import (
	"fmt"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"sync"
)

// nfs相关的操作

type Nfs struct {
	Addr string
	sync.Mutex
	FirstPath string // 预挂载目录
}

// 这个操作实际上是将远程目录 挂载在本地 然后在本地目录创建文件夹 这个创建的文件夹 自然也会同步在远程
func (nfs *Nfs) mount(newPath string) error {
	nfs.Lock()
	defer nfs.Unlock()
	// 这个目录是本地目录
	newerPath := "/temp/" + GetMd516([]byte(newPath))
	_, err := os.Open(newerPath)
	if err != nil {
		if err != os.ErrExist {
			return err
		}
		return nil
	}

	// 在本地目录创建文件夹
	err = os.MkdirAll(newPath, 777)
	if err != nil {
		return err
	}

	// 将nfs挂载到刚刚创建的目录上
	err = exec.Command(fmt.Sprintf("mount -t nfs %s:%s %s", nfs.Addr, "/"+nfs.FirstPath, newerPath)).Start()
	if err != nil {
		klog.Info(err)
		return err
	}
	// 在刚刚那个目录上创建文件夹
	if err = os.MkdirAll(newerPath+"/"+newPath, 777); err != os.ErrExist {
		klog.Info(err)
		return err
	}
	return nil
}

// 需要卸载本地目录 因为目录已经在远程创建好了
func (nfs *Nfs) unmount(path string) {
	newerPath := "/temp/" + GetMd516([]byte(path))
	err := exec.Command(fmt.Sprintf("unmout -v %s", newerPath))
	if err != nil {
		klog.Info(err)
	}
}
