package bitmap

import "sync"

const (
	SHIFT = 6
	MASK  = 0x40
)

type BitMap struct {
	Data []uint64
	Size uint64
	sync.RWMutex
}

func (b *BitMap) Set(value uint64) {
	b.Lock()
	defer b.Unlock()
	index := value >> SHIFT
	location := value % MASK
	b.Data[index%b.Size] = b.Data[index%b.Size] | (1 << location)
}

func (b *BitMap) Exist(value uint64) bool {
	b.RLock()
	defer b.RUnlock()
	index := value >> SHIFT
	location := value & MASK
	if (b.Data[index%b.Size] & (1 << location)) > 0 {
		return true
	}
	return false
}

func (b *BitMap) Clear(value uint64) {
	b.Lock()
	defer b.Unlock()
	index := value >> SHIFT
	location := value & MASK
	b.Data[index%b.Size] = b.Data[index%b.Size] & (^(1 << location))
	return
}

func (b *BitMap) ClearBucket(value uint64) {
	b.Lock()
	defer b.Unlock()
	index := value >> SHIFT
	b.Data[index%b.Size] = 0
}
