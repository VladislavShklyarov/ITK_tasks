package syncPool

import (
	"bytes"
	"sync"
)

var JSONBufferPool = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Поскольку при сериализации используется *bytes.Buffer
// именно его мы и будем переиспользовать
func GetBuffer() *bytes.Buffer {
	return JSONBufferPool.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	JSONBufferPool.Put(buf)
}
