//----------------
//Func  : LRU Container
//Author: xjh
//Date  : 2018/09/10
//Note  :
//----------------
package lru

import (
	"container/list"
	"sync"
)

type node struct {
	k interface{}
	v interface{}
}

type LruContainer struct {
	mutex sync.RWMutex

	maxLen  int // if maxLen<=0 then no limit
	lruList *list.List
	datas   map[interface{}]*list.Element
}

func NewLruContainer(maxLen int) *LruContainer {
	return &LruContainer{
		maxLen:  maxLen,
		lruList: list.New(),
		datas:   make(map[interface{}]*list.Element),
	}
}

//压入队头
func (lc *LruContainer) PushFront(k, v interface{}) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	if e, ok := lc.datas[k]; ok {
		e.Value.(*node).v = v
		lc.lruList.MoveToFront(e)
	} else {
		e := lc.lruList.PushFront(&node{k, v})
		lc.datas[k] = e
	}
	if lc.maxLen > 0 && lc.lruList.Len() > lc.maxLen {
		lc.popBack()
	}
}

//弹出队尾
func (lc *LruContainer) popBack() (interface{}, bool) {
	e := lc.lruList.Back()
	if e != nil {
		lc.del(e.Value.(*node).k)
		return e.Value.(*node).v, true
	}
	return nil, false
}

func (lc *LruContainer) PopBack() (interface{}, bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	return lc.popBack()
}

//查找任意值
func (lc *LruContainer) Get(k interface{}) (interface{}, bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	if e, ok := lc.datas[k]; ok {
		lc.lruList.MoveToFront(e)
		return e.Value.(*node).v, ok
	}
	return nil, false
}

//删除任意值
func (lc *LruContainer) del(k interface{}) {
	if e, ok := lc.datas[k]; ok {
		delete(lc.datas, k)
		lc.lruList.Remove(e)
	}
}

//删除任意值
func (lc *LruContainer) Del(k interface{}) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	lc.del(k)
}

//长度
func (lc *LruContainer) Len() int {
	lc.mutex.RLock()
	defer lc.mutex.RUnlock()
	return lc.lruList.Len()
}
