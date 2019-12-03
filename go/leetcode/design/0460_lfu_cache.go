package design

import (
	"container/heap"
	"time"
)

type entry struct {
	key, value int
	frequence, index int
	date time.Time
}

type PQ []*entry

func (pq PQ) Len() int {return len(pq)}

func (pq PQ) Less(i, j int) bool {
	if pq[i].frequence == pq[j].frequence {
		return pq[i].date.Before(pq[j].date)
	}
	return pq[i].frequence < pq[j].frequence
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{})  {
	n := len(*pq)
	entry := x.(*entry)
	entry.index = n
	entry.date = time.Now()
	*pq = append(*pq, entry)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*pq = old[0:n-1]
	return entry
}

func (pq *PQ) update(entry *entry) {
	entry.frequence++
	entry.date = time.Now()
	heap.Fix(pq, entry.index)
}

type LFUCache struct {
	m map[int]*entry
	pq PQ
	capacity int
}

func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		m: make(map[int]*entry, capacity),
		pq: make(PQ, 0, capacity),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	ep, ok := this.m[key]
	if ok {
		this.pq.update(ep)
		return ep.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	ep, ok := this.m[key]
	if ok {
		this.m[key].value = value
		this.pq.update(ep)
		return
	}

	ep = &entry{key: key, value: value}
	if len(this.pq) == this.capacity {
		temp := heap.Pop(&this.pq).(*entry)
		delete(this.m, temp.key)
	}
	this.m[key] = ep
	heap.Push(&this.pq, ep)
}