package ebus

import (
	"context"
	"sync"
)

const (
	Sync    = true
	NotSync = false
)

var eventMgr *EvnetMgr

// 发布事件
func Publish(ctx context.Context, topic string, arg any, sync bool) {
	eventMgr.mux.Lock()
	defer eventMgr.mux.Unlock()
	data, ok := eventMgr.topic[topic]
	if !ok {
		return
	}
	if sync {
		for _, v := range data {
			v.Handler(ctx, arg)
		}
	} else {
		for _, v := range data {
			go v.Handler(ctx, arg)
		}
	}
}

// 订阅事件
func Subscribe(topic string, l ...*Listener) {
	eventMgr.mux.Lock()
	defer eventMgr.mux.Unlock()
	eventMgr.topic[topic] = append(eventMgr.topic[topic], l...)
}

// 删除事件订阅
func Remove(topic string, l *Listener) {
	eventMgr.mux.Lock()
	defer eventMgr.mux.Unlock()
	data, ok := eventMgr.topic[topic]
	if !ok {
		return
	}

	for i, v := range data {
		if v == l {
			eventMgr.topic[topic] = append(eventMgr.topic[topic][:i], eventMgr.topic[topic][i+1:]...)
			return
		}
	}
}

// 事件处理器
type Handler func(ctx context.Context, arg any)

// 监听器
type Listener struct {
	Handler Handler
}

func NewListener(f Handler) *Listener {
	return &Listener{
		Handler: f,
	}
}

// 事件管理器
type EvnetMgr struct {
	mux   sync.Mutex
	topic map[string][]*Listener
}

func NewMgr() {
	eventMgr = &EvnetMgr{
		mux:   sync.Mutex{},
		topic: make(map[string][]*Listener),
	}
}
