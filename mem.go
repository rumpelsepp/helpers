package helpers

import "sync"

func CreateMemPool() sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return make([]byte, 4*1024)
		},
	}
}

func GetSlice(pool *sync.Pool, size int) []byte {
	buf := pool.Get().([]byte)
	if len(buf) > size {
		buf = buf[:size]
	} else if len(buf) < size {
		buf = make([]byte, size)
	}
	return buf
}
