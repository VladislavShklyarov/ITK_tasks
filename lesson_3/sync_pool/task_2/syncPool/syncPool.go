package syncPool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type RequestData struct {
	Data map[string]string `json:"data"`
}

func (rd *RequestData) Reset() {
	for key := range rd.Data {
		delete(rd.Data, key)
	}
}

func NewRequestData() *RequestData {
	atomic.AddInt64(&Created, 1)
	return &RequestData{
		Data: make(map[string]string),
	}
}

var Created int64
var requestDataPool = &sync.Pool{
	New: func() interface{} {
		atomic.AddInt64(&Created, 1)
		return NewRequestData()
	},
}

type SyncPool struct{}

func NewSyncPool() *SyncPool {
	return &SyncPool{}
}

func (p *SyncPool) Get() *RequestData {
	return requestDataPool.Get().(*RequestData)
}

func (p *SyncPool) Put(rd *RequestData) {
	rd.Reset()
	fmt.Printf("req ptr: %p\n", rd.Data)

	requestDataPool.Put(rd)
}
