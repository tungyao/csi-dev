package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"os"
	"sync"
	"unsafe"
)

const (
	// MaxSpaceCol 最大同时支持最大的并发数
	MaxSpaceCol = 1<<4 - 1
)
const (
	SpaceDirDeleting = iota << 1
	SpaceDirPre
	SpaceDirDeleted
)

// 数据存储的具体实现方案
// 后续可以有很大的提升空间

type LocalStorage struct {
	sync.RWMutex
	space map[string]*Space
}

type Space struct {
	ctx      context.Context
	Name     string
	Operator chan int
	ls       unsafe.Pointer
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

// NewLocalStorageHub TODO 监听
func (ls *LocalStorage) NewLocalStorageHub() {

}

// NewSpace 开辟新的一个空间
func (ls *LocalStorage) NewSpace() (*Space, context.Context) {
	sp := &Space{}
	sp.Name = uuid.New().String()
	sp.Operator = make(chan int, MaxSpaceCol)
	sp.ctx = context.Background()
	ls.RLock()
	ls.space[sp.Name] = sp
	os.Mkdir(sp.Name, 777)
	sp.ls = unsafe.Pointer(ls)
	ls.Unlock()
	return sp, sp.ctx
}

func (ls *LocalStorage) GetSpace(spaceName string) (*Space, error) {
	ls.Lock()
	defer ls.Unlock()
	sp, ok := ls.space[spaceName]
	if !ok {
		return nil, errors.New("no space")
	}
	return sp, nil
}

// Free 释放该文件夹
func (sp *Space) Free() {
	ls := (*LocalStorage)(sp.ls)
	ls.Lock()
	sp.Operator <- SpaceDirDeleting
	defer ls.Unlock()
	os.Remove(sp.Name)
	sp.Operator <- SpaceDirDeleted
}
