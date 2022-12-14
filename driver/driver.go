package driver

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

// nfs相关的操作

type Nfs struct {
	Addr string
	sync.Mutex
}

func (nfs *Nfs) mount(newPath string) error {
	nfs.Lock()
	defer nfs.Unlock()
	_, err := os.Open("/temp/" + newPath)
	if err != nil {
		if err != os.ErrExist {
			return err
		}
		return nil
	}
	newPath = "/temp/" + newPath
	err = os.MkdirAll(newPath, 655)
	if err != nil {
		return err
	}
	exec.Command(fmt.Sprintf("mount -t nfs %s:%s %s", nfs.Addr, "/", newPath))
	// 新建目录
	return nil
}
func (nfs *Nfs) unmount(path string) {

}
