package main

import (
	"github.com/google/uuid"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
)

type Nfs struct {
	ip       string
	bashPath string
}

type NfsDt struct {
	localPath  string
	remotePath string
	err        error
	name       string
}

func NewNfs(link string, bashPath string) *Nfs {
	return &Nfs{ip: link, bashPath: bashPath}
}

// 将远程目录挂载到本地指定目录
func (n *Nfs) mount(remote, local string) {

	err := exec.Command("mount", "-t", "nfs", n.ip+":"+n.bashPath+remote, local).Run()
	if err != nil {
		klog.Errorln("挂载目录错误", n.ip, n.bashPath, local)
		return
	}
}

// 将目录挂载到本地一个临时目录
func (n *Nfs) provisionalPath(ud ...string) *NfsDt {
	var localRandom string
	if len(ud) > 0 {
		localRandom = ud[0]
	} else {
		localRandom = uuid.New().String()
	}

	err := os.Mkdir("/mnt/"+localRandom, 777)
	if err != nil {
		return &NfsDt{
			err: err,
		}
	}
	n.mount("", "/mnt/"+localRandom)
	return &NfsDt{
		localPath:  "/mnt/" + localRandom,
		remotePath: "",
		name:       localRandom,
	}
}

func (n *Nfs) umount(path string) error {
	err := exec.Command("umount", path).Run()
	if err != nil {
		klog.Errorln("删除目录错误", err)
		return err
	}
	return nil
}

// 在远程创建一个目录
func (n *NfsDt) createPath(path string) {
	err := os.Mkdir(n.localPath+path, 777)
	if err != nil {
		klog.Errorln("创建目录错误", path)
		return
	}
}

// 在远程删除一个目录
func (n *NfsDt) deletePath() error {
	return nil
}
