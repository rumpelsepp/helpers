package helpers

import "sync"

type Broadcaster struct {
	InCh    chan []byte
	OutChs  []chan []byte
	MemPool *sync.Pool
	WG      *sync.WaitGroup
}

func (bc *Broadcaster) Serve() {
	for data := range bc.InCh {
		for _, listener := range bc.OutChs {
			buf := GetSlice(bc.MemPool, len(data))
			n := copy(buf, data)
			listener <- buf[:n]
		}
	}
	for _, ch := range bc.OutChs {
		close(ch)
	}
	bc.WG.Done()
}
