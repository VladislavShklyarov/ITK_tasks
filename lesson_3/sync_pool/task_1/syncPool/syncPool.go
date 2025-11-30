package syncPool

import "sync"

const Capacity int = 1024

type SyncPool struct{}

func NewSyncPool() *SyncPool {
	return &SyncPool{}
}

var BuffersPool = &sync.Pool{
	New: func() interface{} {
		b := make([]byte, 0, Capacity) // поставим объем буффера в 1024 байта.
		return &b
	},
}

func (p *SyncPool) GetBytes() *[]byte {
	return BuffersPool.Get().(*[]byte)
}

func (p *SyncPool) PutBytes(b *[]byte) {
	if cap(*b) > Capacity {
		return
	}
	*b = (*b)[:0]
	BuffersPool.Put(b)
}
