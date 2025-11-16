package cache

import (
	"sync/atomic"
	"time"
)

// 缓存命中率
type CacheStat struct {
	hitNum       uint64
	missNum      uint64
	statInterval time.Duration
	sizeCallback func(hit uint64, miss uint64)
}

func NewCacheStat(statInterval time.Duration, sizeCallback func(hit uint64, miss uint64)) *CacheStat {
	st := &CacheStat{
		statInterval: statInterval,
		sizeCallback: sizeCallback,
	}
	go st.statLoop()
	return st
}

func (c *CacheStat) statLoop() {
	ticker := time.NewTicker(c.statInterval)
	defer ticker.Stop()

	for range ticker.C {
		hit := atomic.SwapUint64(&c.hitNum, 0)
		miss := atomic.SwapUint64(&c.missNum, 0)
		c.sizeCallback(hit, miss)
	}
}

// 命中
func (c *CacheStat) IncrementHit() {
	atomic.AddUint64(&c.hitNum, 1)
}

// 未命中
func (c *CacheStat) IncrementMiss() {
	atomic.AddUint64(&c.missNum, 1)
}
